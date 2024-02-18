package atq

import (
	"context"
	"time"

//	"github.com/valteem/atq/internal/base"
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

type TaskInfo struct {
	ID string // task identifier (missing in Task definition?!)
	Qeueu string
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