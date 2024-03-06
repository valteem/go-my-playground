package context

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
	"github.com/valteem/atq/internal/base"
)

func TestCreateContextWithFutureDeadline(t *testing.T) {
	tests := []struct {
		deadline time.Time
	}{
		{
			time.Now().Add(time.Hour),
		},
	}

	for _, tst := range tests {
		msg := &base.TaskMessage{
			Type:    "some type",
			ID:      uuid.NewString(),
			Payload: nil,
		}

		ctx, cancel := New(context.Background(), msg, tst.deadline)
		select {
		case x := <-ctx.Done():
			t.Errorf("<-ctx.Done() == %+v, expected nothing (must block)", x) // select before cancel()
		default:
		}

		result, ok := ctx.Deadline()
		if !ok {
			t.Errorf("ctx.Deadline() returned false, expect true as deadline has been set")
		}
		if !cmp.Equal(tst.deadline, result) {
			t.Errorf("ctx.Deadline() returned %+v, expect %+v", result, tst.deadline)
		}

		cancel()

		select {
		case <-ctx.Done():
		default:
			t.Errorf("ctx.Done() blocked, expected to be non-blocking")
		}
	}
}

func TestCreateContextWithPastDeadline(t *testing.T) {
	tests := []struct {
		deadline time.Time
	}{
		{
			time.Now().Add(-1 * time.Hour),
		},
	}
	for _, tst := range tests {
		msg := &base.TaskMessage{
			Type:    "someType",
			ID:      uuid.NewString(),
			Payload: nil,
		}

		ctx, cancel := New(context.Background(), msg, tst.deadline)
		defer cancel()

		select {
		case <-ctx.Done():
		default:
			t.Errorf("cancel signal not expected")
		}

		result, ok := ctx.Deadline()
		if !ok {
			t.Errorf("deadline is set, no false return value expected")
		}
		if !cmp.Equal(tst.deadline, result) {
			t.Errorf("ctx.Deadline() returned %+v, expected %+v", result, tst.deadline)
		}

	}
}

func TestCreateContextWithBaseContext(t *testing.T) {
	// https://go.dev/blog/context
	type ctxKey string
	type ctxValue string
	var key ctxKey = "key"
	var value ctxValue = "value"

	tests := []struct {
		baseCtx  context.Context
		validate func(ctx context.Context, t *testing.T) error
	}{
		{
			baseCtx: context.WithValue(context.Background(), key, value), // passing key and value of custom types to avoid collision
			validate: func(ctx context.Context, t *testing.T) error {
				result, ok := ctx.Value(key).(ctxValue)
				if !ok {
					return fmt.Errorf("ctx.Value() returned false, expected true")
				}
				if expected := value; result != expected {
					return fmt.Errorf("ctxValue() returned %+v, expect %s", result, expected)
				}
				return nil
			},
		},
	}

	for _, tst := range tests {
		msg := &base.TaskMessage{
			Type:    "someType",
			ID:      uuid.NewString(),
			Payload: nil,
		}

		ctx, cancel := New(tst.baseCtx, msg, time.Now().Add(30*time.Minute))
		defer cancel()

		select {
		case x := <-ctx.Done():
			t.Errorf("<-ctx.Done() returned %+v, expect nothing", x)
		default:
		}

		if e := tst.validate(ctx, t); e != nil {
			t.Errorf("%+v", e)
		}
	}
}

func TestGetTaskMetadataFromContext(t *testing.T) {
	tests := []struct {
		description string
		msg         *base.TaskMessage
	}{
		{
			description: "zero retry count",
			msg: &base.TaskMessage{
				Type:    "defaultType",
				ID:      uuid.NewString(),
				Retry:   100,
				Retried: 0,
				Timeout: 3600,
				Queue:   "default",
			},
		},
	}
	for _, tst := range tests {
		ctx, cancel := New(context.Background(), tst.msg, time.Now().Add(60*time.Minute))
		defer cancel()
		id, ok := GetTaskID(ctx)
		if !ok {
			t.Errorf("%s: GetTaskID() returned false", tst.description)
		}
		if ok && id != tst.msg.ID {
			t.Errorf("%s: GetsTaskID() returned %q, expected %q", tst.description, id, tst.msg.ID)
		}
		retried, ok := GetRetryCount(ctx)
		if !ok {
			t.Errorf("%s: GetRetryCount() returned false", tst.description)
		}
		if ok && retried != tst.msg.Retried {
			t.Errorf("%s: GetRetryCount() returned %q, expected %q", tst.description, retried, tst.msg.Retried)
		}
		qname, ok := GetQueueName(ctx)
		if !ok {
			t.Errorf("%s: GetQueueName() returned false", tst.description)
		}
		if ok && qname != tst.msg.Queue {
			t.Errorf("%s: GetsQueueName() returned %q, expected %q", tst.description, qname, tst.msg.Queue)
		}
	}
}
