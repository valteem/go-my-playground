package base

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Internal representation of a task, with additional metadata
type TaskMessage struct {
	Type         string
	Payload      []byte
	ID           string
	Queue        string
	Retry        int
	Retried      int
	ErrorMsg     string
	LastFailedAt int64 // Unix time format, zero if no failures
	Timeout      int64
	Deadline     int64
	UniqueKey    string // redis key for uniqueness lock, empty if no lock was used
	GroupKey     string // TaskInfo.Group ? empty if no aggregation is planned for the task
	Retention    int64  // how long (in seconds) retain the task after completion
	CompletedAt  int64
}

// sorted set member (whatever this may mean)
type Z struct {
	Message *TaskMessage
	Score   int64
}

type ServerInfo struct {
	Host              string
	PID               int
	ServerID          string
	Concurrency       int
	Queues            map[string]int
	StrictPriority    bool
	Status            string
	Started           time.Time
	ActiveWorkerCount int
}

type WorkerInfo struct {
	Host     string
	PID      int
	ServerID string
	ID       string
	Type     string
	Payload  []byte
	Queue    string
	Started  time.Time
	Deadline time.Time
}

// Message broker - manages task queues
type Broker interface {
	// General
	Ping() error
	Close() error
	Enqueue(ctx context.Context, msg *TaskMessage) error
	EnqueueUnique(ctx context.Context, msg *TaskMessage, ttl time.Duration) error
	Dequeue(qnames ...string) (*TaskMessage, time.Time, error)
	Done(ctx context.Context, msg *TaskMessage) error
	MarkAsComplete(ctx context.Context, msg *TaskMessage) error
	Requeue(ctx context.Context, msg *TaskMessage) error
	Schedule(ctx context.Context, msg *TaskMessage, processAt time.Time) error
	ScheduleUnique(ctx context.Context, msg *TaskMessage, processAt time.Time, ttl time.Duration) error
	Retry(ctx context.Context, msg *TaskMessage, processAt time.Time, errMsg string, isFailure bool) error
	Archive(ctx context.Context, msg *TaskMessage, errMsg string) error
	ForwardIfReady(qnames ...string) error
	// Group aggregation
	AddToGroup(ctx context.Context, msg *TaskMessage, qname string) error
	AddToGroupUnique(ctx context.Context, msg *TaskMessage, groupKey string, ttl time.Duration) error
	ListGroups(qname string) ([]string, error)
	AggregationCheck(qname, gname string, t time.Time, gracePeriod, maxDelay time.Duration, maxSize int) (aggregationSetID string, err error)
	ReadAggregationSet(qname, gname, aggregationSetID string) ([]*TaskMessage, time.Time, error)
	DeleteAggregationSet(ctx context.Context, qname, gname, aggregationSetID string) error
	ReclaimStaleAggregationSets(qname string) error
	// Task retention
	DeleteExpiredCompletedtasks(qname string) error
	// Lease (whatever this may mean)
	ListLeaseExpired(cutoff time.Time, qnames ...string) ([]*TaskMessage, error)
	Extendlease(qname string, ids ...string) (time.Time, error)
	// State snapshot
	WriteServerState(info *ServerInfo, workers []*WorkerInfo, ttl time.Duration) error
	ClearServerState(host string, pid int, serverID string) error
	// Cancelation
	CanceletionPubSub() (*redis.PubSub, error)
	PublishCancelation(id string) error
	// Write (?)
	WriteResult(qname, id string, data []byte) (n int, err error)
}
