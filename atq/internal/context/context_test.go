package context

import (
	"context"
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