package testutil

import (
	"time"

	"github.com/google/uuid"
	"github.com/valteem/atq/internal/base"
)

func makeDefaultTaskMessage() *base.TaskMessage {
	return &base.TaskMessage{
		ID:       uuid.NewString(),
		Type:     "default_task",
		Queue:    "default",
		Retry:    50, // because why not
		Timeout:  3600,
		Deadline: 0, // no deadline
	}
}

type TaskMessageBuilder struct {
	msg *base.TaskMessage
}

func NewTaskMessageBuilder() *TaskMessageBuilder {
	return &TaskMessageBuilder{}
}

func (b *TaskMessageBuilder) lazyInit() { // why 'lazy' ?
	if b.msg == nil { // because it initializes TaskMessage only if it is initialized yet
		b.msg = makeDefaultTaskMessage()
	}
}

func (b *TaskMessageBuilder) Build() *base.TaskMessage {
	b.lazyInit()
	return b.msg
}

func (b *TaskMessageBuilder) SetType(typename string) *TaskMessageBuilder {
	b.lazyInit()
	b.msg.Type = typename
	return b
}

func (b *TaskMessageBuilder) SetPayload(payload []byte) *TaskMessageBuilder {
	b.lazyInit()
	b.msg.Payload = payload
	return b
}

func (b *TaskMessageBuilder) SetQueue(queue string) *TaskMessageBuilder {
	b.lazyInit()
	b.msg.Queue = queue
	return b
}

func (b *TaskMessageBuilder) SetRetry(retry int) *TaskMessageBuilder {
	b.lazyInit()
	b.msg.Retry = retry
	return b
}

func (b *TaskMessageBuilder) SetTimeout(timeout time.Duration) *TaskMessageBuilder {
	b.lazyInit()
	b.msg.Timeout = int64(timeout.Seconds())
	return b
}

func (b *TaskMessageBuilder) SetDeadLine(deadline time.Time) *TaskMessageBuilder {
	b.lazyInit()
	b.msg.Deadline = deadline.Unix()
	return b
}

func (b *TaskMessageBuilder) SetGroup(groupkey string) *TaskMessageBuilder {
	b.lazyInit()
	b.msg.GroupKey = groupkey
	return b
}
