package atq

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/valteem/atq/internal/base"
)

type ResultWriter struct {
	id     string // task ID
	qname  string // qeueu name this task belongs to
	broker any // base.Broker /internal/base
	// TODO: find SO thread about packing context inside struct considered a bad idea
	ctx    context.Context
}

type Task struct {
	typename string // type of task
	payload  []byte // data needed to perform the task
	opts     []Option
	w        *ResultWriter
}

func (t *Task) Type() string                { return t.typename }
func (t *Task) Payload() []byte             { return t.payload  }
func (t *Task) ResultWriter() *ResultWriter {return t.w} // nil ptr on newly created task

// Public constructor: accepts task type name, payload, options
func NewTask(typename string, payload[]byte, opts ...Option) *Task {
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
	TaskStateActive: "active",
	TaskStatePending: "pending",
	TaskStateScheduled: "scheduled",
	TaskStateRetry: "retry",
	TaskStateArchived: "archived",
	TaskStateCompleted: "completed",
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
	ID string // task identifier (missing in Task definition?!)
	Queue string
	Type string // typename?
	Payload []byte
	State TaskState
	MaxRetry int
	Retried int
	LastErr string // error message from last failure
	LastFailedAt time.Time
	Timeout time.Duration // time before being retried
	Deadline time.Time // zero if not set
	Group string
	NextProcessAt time.Time // scheduled processing time, zero if not applicable
	IsOrphaned bool // active task with no worker processing it (worker crash, network failure, etc)
	Retention time.Duration // time after task is successfully processed
	CompletedAt time.Time
	Result []byte
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
		ID: msg.ID,
		Queue: msg.Queue,
		Type: msg.Type,
		Payload: msg.Payload,
		MaxRetry: msg.Retry,
		Retried: msg.Retried,
		LastErr: msg.ErrorMsg,
		Group: msg.GroupKey,
		Timeout: time.Duration(msg.Timeout) * time.Second,
		Deadline: fromUnixTimeOrZero(msg.Deadline),
		Retention: time.Duration(msg.Retention) * time.Second,
		NextProcessAt: nextProcessAt,
		LastFailedAt: fromUnixTimeOrZero(msg.LastFailedAt),
		CompletedAt: fromUnixTimeOrZero(msg.CompletedAt),
		Result: result,
		// TODO: replace with switch-case if base.TaskState is used instead
		State: state,
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
	Network string // default TCP
	Addr string // host:port
	Username string
	Password string
	DB int // Redis DB to select after connecting to a server
	DialTimeout time.Duration // establish new connection, default 5 sec
	ReadTimeout time.Duration // socket reads, default 3 sec
	WriteTimeout time.Duration //socket writes, default 3 sec
	PoolSize int // max number of socket connections, default 10 per CPU
	TLSConfig *tls.Config // TLS negotiated only if this field is set
}