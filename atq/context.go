// Looks like a wrapper around internal context

package atq

import (
	"context"

	atqctx "github.com/valteem/atq/internal/context"
)

func GetTaskID(ctx context.Context) (id string, ok bool) {
	return atqctx.GetTaskID(ctx)
}

func GetRetryCount(ctx context.Context) (n int, ok bool) {
	return atqctx.GetRetryCount(ctx)
}

func GetMaxRetry(ctx context.Context) (n int, ok bool) {
	return atqctx.GetMaxRetry(ctx)
}

func GetQueueName(ctx context.Context) (queue string, ok bool) {
	return atqctx.GetQueueName(ctx)
}
