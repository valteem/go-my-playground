package atq

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/valteem/atq/internal/base"
)

type ResultWriter struct {
	id     string          // task ID
	qname  string          // qeueu name this task belongs to
	broker any             // base.Broker /internal/base
	ctx    context.Context // TODO: find SO thread about packing context inside struct considered a bad idea
}

type Task struct {
	typename string // type of task
	payload  []byte // data needed to perform the task
	opts     []Option
	w        *ResultWriter
}

func (t *Task) Type() string                { return t.typename }
func (t *Task) Payload() []byte             { return t.payload }
func (t *Task) ResultWriter() *ResultWriter { return t.w } // nil ptr on newly created task

// Public constructor: accepts task type name, payload, options
func NewTask(typename string, payload []byte, opts ...Option) *Task {
	return &Task{
		typename: typename,
		payload:  payload,
		opts:     opts,
	}
}

// Private constructor: accepts task type name, payload, result writer
func newTask(typename string, payload []byte, w *ResultWriter) *Task {
	return &Task{
		typename: typename,
		payload:  payload,
		w:        w,
	}
}

type TaskState int

const (
	TaskStateActive TaskState = iota + 1
	TaskStatePending
	TaskStateScheduled
	TaskStateRetry
	TaskStateArchived
	TaskStateCompleted
	TaskStateAggregating // waiting in a group to be aggregated into single task
)

var taskStateString = map[TaskState]string{
	TaskStateActive:      "active",
	TaskStatePending:     "pending",
	TaskStateScheduled:   "scheduled",
	TaskStateRetry:       "retry",
	TaskStateArchived:    "archived",
	TaskStateCompleted:   "completed",
	TaskStateAggregating: "aggregating",
}

// Use map instead of switch-case (however switch is 3x faster https://stackoverflow.com/a/73275004)
func (s TaskState) String() string {
	str, ok := taskStateString[s]
	if !ok {
		panic("atq: unknown task state")
	}
	return str
}

type TaskInfo struct {
	ID            string // task identifier (missing in Task definition?!)
	Queue         string
	Type          string // typename?
	Payload       []byte
	State         TaskState
	MaxRetry      int
	Retried       int
	LastErr       string // error message from last failure
	LastFailedAt  time.Time
	Timeout       time.Duration // time before being retried
	Deadline      time.Time     // zero if not set
	Group         string
	NextProcessAt time.Time     // scheduled processing time, zero if not applicable
	IsOrphaned    bool          // active task with no worker processing it (worker crash, network failure, etc)
	Retention     time.Duration // time after task is successfully processed
	CompletedAt   time.Time
	Result        []byte
}

// Convert Unix time (int64 in seconds since 1970-1-1 00:00:00) to time.Time
func fromUnixTimeOrZero(t int64) time.Time {
	if t == 0 {
		return time.Time{} // 'real' zero time (time.Time(0001-01-01T00:00:00Z))
		// instead of Unix zero time (time.Time(1970-01-01T00:00:00Z))
	}
	return time.Unix(t, 0)
}

// TODO: replace (if really needed) TaskSate with base.TaskState, both look the same
func newTaskInfo(msg base.TaskMessage, state TaskState, nextProcessAt time.Time, result []byte) *TaskInfo {
	info := TaskInfo{
		ID:            msg.ID,
		Queue:         msg.Queue,
		Type:          msg.Type,
		Payload:       msg.Payload,
		MaxRetry:      msg.Retry,
		Retried:       msg.Retried,
		LastErr:       msg.ErrorMsg,
		Group:         msg.GroupKey,
		Timeout:       time.Duration(msg.Timeout) * time.Second,
		Deadline:      fromUnixTimeOrZero(msg.Deadline),
		Retention:     time.Duration(msg.Retention) * time.Second,
		NextProcessAt: nextProcessAt,
		LastFailedAt:  fromUnixTimeOrZero(msg.LastFailedAt),
		CompletedAt:   fromUnixTimeOrZero(msg.CompletedAt),
		Result:        result,
		State:         state, // TODO: replace with switch-case if base.TaskState is used instead
	}
	return &info
}

// 'Discriminated union of types':
// - RedisClientOpt
// - RedisFailoverClientOpt
// - RedisClusterClientOpt
type RedisConnOpt interface {
	MakeRedisClient() any
}

type RedisClientOpt struct {
	Network      string // default TCP
	Addr         string // host:port
	Username     string
	Password     string
	DB           int           // Redis DB to select after connecting to a server
	DialTimeout  time.Duration // establish new connection, default 5 sec
	ReadTimeout  time.Duration // socket reads, default 3 sec
	WriteTimeout time.Duration //socket writes, default 3 sec
	PoolSize     int           // max number of socket connections, default 10 per CPU
	TLSConfig    *tls.Config   // TLS negotiated only if this field is set
}

func (opt RedisClientOpt) MakeRedisClient() any {
	return redis.NewClient(&redis.Options{
		Network:      opt.Network,
		Addr:         opt.Addr,
		Username:     opt.Username,
		Password:     opt.Password,
		DB:           opt.DB,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		PoolSize:     opt.PoolSize,
		TLSConfig:    opt.TLSConfig,
	})
}

type RedisFailoverClientOpt struct {
	MasterName       string   // 'master' monitored by 'sentinels'
	SentinelAddrs    []string // host:port
	SentinelPassword string
	Username         string
	Password         string
	DB               int
	DialTimeout      time.Duration
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	PoolSize         int
	TLSConfig        *tls.Config
}

func (opt RedisFailoverClientOpt) MakeRedisClient() any {
	return redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       opt.MasterName,
		SentinelAddrs:    opt.SentinelAddrs,
		SentinelPassword: opt.SentinelPassword,
		Username:         opt.Username,
		Password:         opt.Password,
		DB:               opt.DB,
		DialTimeout:      opt.DialTimeout,
		ReadTimeout:      opt.ReadTimeout,
		WriteTimeout:     opt.WriteTimeout,
		PoolSize:         opt.PoolSize,
		TLSConfig:        opt.TLSConfig,
	})
}

type RedisClusterClientOpt struct {
	Addrs        []string // host:port addresses of cluster node
	MaxRedirects int
	Username     string
	Password     string
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TLSConfig    *tls.Config
}

func (opt RedisClusterClientOpt) MakeRedisClient() any {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:        opt.Addrs,
		MaxRedirects: opt.MaxRedirects,
		Username:     opt.Username,
		Password:     opt.Password,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
		TLSConfig:    opt.TLSConfig,
	})
}

func parseRedisURI(u *url.URL) (RedisConnOpt, error) {
	var db int
	var err error
	var redisConnOpt RedisClientOpt

	if len(u.Path) > 0 {
		s := strings.Split(strings.Trim(u.Path, "/"), "/") // remove leading and trailing slashes, split to segments
		db, err = strconv.Atoi(s[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing redis uri: db number must be first segment of the path")
		}
	}
	var password string
	if v, ok := u.User.Password(); ok { // URL.Password() returns URL.password, URL.passwordSet
		password = v
	}

	if u.Scheme == "rediss" { // 'redis' vs 'rediss' same as 'http' vs 'https'
		h, _, err := net.SplitHostPort(u.Host)
		if err != nil {
			h = u.Host
		}
		redisConnOpt.TLSConfig = &tls.Config{ServerName: h}
	}

	redisConnOpt.Addr = u.Host
	redisConnOpt.Password = password
	redisConnOpt.DB = db

	return redisConnOpt, nil
}

func parseRedisSocketUri(u *url.URL) (RedisConnOpt, error) {
	errPrefix := "atq: error parsing redis socket uri"
	if len(u.Path) == 0 {
		return nil, fmt.Errorf("%s: path does not exist", errPrefix)
	}
	q := u.Query()
	var db int
	var err error
	if n := q.Get("db"); n != "" {
		db, err = strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("%s: query parameter 'db' must be a number", errPrefix)
		}
	}
	var password string
	if v, ok := u.User.Password(); ok {
		password = v
	}
	return RedisClientOpt{Network: "unix", Addr: u.Path, DB: db, Password: password}, nil
}
