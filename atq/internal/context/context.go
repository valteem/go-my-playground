package context

import (
	"context"
	"time"

	"github.com/valteem/atq/internal/base"
)

// Task data to be added to context
type taskMetadata struct {
	id         string
	maxRetry   int
	retryCount int
	qname      string
}

type ctxKey int // unexported to prevent collisions outside the package

const metadataCtxKey ctxKey = 0 // arbitrary value

func New(base context.Context, msg *base.TaskMessage, deadline time.Time) (context.Context, context.CancelFunc) {
	metadata := taskMetadata{
		id:         msg.ID,
		maxRetry:   msg.Retry,
		retryCount: msg.Retried,
		qname:      msg.Queue,
	}
	return context.WithDeadline(context.WithValue(base, metadataCtxKey, metadata), deadline)
}

func GetTaskID(ctx context.Context) (string, bool) {
	metadata, ok := ctx.Value(metadataCtxKey).(taskMetadata)
	if !ok {
		return "", false
	}
	return metadata.id, true
}

func GetRetryCount(ctx context.Context) (int, bool) {
	metadata, ok := ctx.Value(metadataCtxKey).(taskMetadata)
	if !ok {
		return 0, false
	}
	return metadata.retryCount, true
}

func GetMaxRetry(ctx context.Context) (int, bool) {
	metadata, ok := ctx.Value(metadataCtxKey).(taskMetadata)
	if !ok {
		return 0, false
	}
	return metadata.maxRetry, true
}

func GetQueueName(ctx context.Context) (string, bool) {
	metadata, ok := ctx.Value(metadataCtxKey).(taskMetadata)
	if !ok {
		return "", false
	}
	return metadata.qname, true
}
