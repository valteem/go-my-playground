package testutil

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/valteem/atq/internal/base"
)

func TestTaskMessageBuilder(t *testing.T) {

	tests := []struct {
		description string
		op          func(b *TaskMessageBuilder) // operation to perform on Task Builder
		expect      *base.TaskMessage
	}{
		{
			description: "default build",
			op:          nil,
			expect: &base.TaskMessage{
				Type:     "default_task",
				Queue:    "default",
				Payload:  nil,
				Retry:    50,
				Timeout:  3600,
				Deadline: 0,
			},
		},
		{
			description: "custom type, payload, queue",
			op: func(b *TaskMessageBuilder) {
				b.SetType("custom_type").SetPayload([]byte("message")).SetQueue("newqueue")
			},
			expect: &base.TaskMessage{
				Type:     "custom_type",
				Queue:    "newqueue",
				Payload:  []byte("message"),
				Retry:    50,
				Timeout:  3600,
				Deadline: 0,
			},
		},
		{
			description: "custom retry, timeout, deadline",
			op: func(b *TaskMessageBuilder) {
				b.SetRetry(1).
					SetTimeout(30 * time.Second).
					SetDeadLine(time.Date(2024, 12, 31, 23, 59, 59, 0, time.FixedZone("UTC+7", 7*60*60)))
			},
			expect: &base.TaskMessage{
				Type:     "default_task",
				Queue:    "default",
				Payload:  nil,
				Retry:    1,
				Timeout:  30,
				Deadline: time.Date(2024, 12, 31, 23, 59, 59, 0, time.FixedZone("UTC+7", 7*60*60)).Unix(),
			},
		},
		{
			description: "custom group",
			op: func(b *TaskMessageBuilder) {
				b.SetGroup("newgroup")
			},
			expect: &base.TaskMessage{
				Type:     "default_task",
				Queue:    "default",
				Payload:  nil,
				Retry:    50,
				Timeout:  3600,
				Deadline: 0,
				GroupKey: "newgroup",
			},
		},
	}

	cmpOpts := []cmp.Option{cmpopts.IgnoreFields(base.TaskMessage{}, "ID")}

	for _, tst := range tests {
		var b TaskMessageBuilder
		if tst.op != nil {
			tst.op(&b)
		}

		result := b.Build()
		if diff := cmp.Diff(tst.expect, result, cmpOpts...); diff != "" {
			t.Errorf("%s, TaskMessageBuilder,Build() = %+v, expect %+v\n(-expect.+result)\n%s",
				tst.description, result, tst.expect, diff)
		}
	}
}
