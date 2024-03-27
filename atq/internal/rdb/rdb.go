package rdb

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/valteem/atq/internal/errors"
	"github.com/valteem/atq/internal/timeutil"
)

const statsTTL = 90 * 24 * time.Hour // https://redis.io/commands/ttl/ (?)

const LeaseDuration = 30 * time.Second

// Client interface to query and mutate task queues
type RDB struct {
	client redis.UniversalClient
	clock  timeutil.Clock
}

func NewRDB(client redis.UniversalClient) *RDB {
	return &RDB{
		client: client,
		clock:  timeutil.NewRealClock(),
	}
}

func (r *RDB) Close() error {
	return r.client.Close()
}

func (r *RDB) Client() redis.UniversalClient {
	return r.client
}

func (r *RDB) SetClock(c timeutil.Clock) {
	r.clock = c
}

func (r *RDB) Ping() error {
	return r.client.Ping(context.Background()).Err()
}

func (r *RDB) runScript(ctx context.Context, op errors.Op, script *redis.Script, keys []string, args ...any) error {
	if err := script.Run(ctx, r.client, keys, args...).Err(); err != nil {
		return errors.E(op, errors.Internal, fmt.Sprintf("redis eval error: %v", err))
	}
	return nil
}

func (r *RDB) runScriptWithErrorCode(ctx context.Context, op errors.Op, script *redis.Script, keys []string, args ...any) (int64, error) {
	res, err := script.Run(ctx, r.client, keys, args...).Result()
	if err != nil {
		return 0, errors.E(op, errors.Unknown, fmt.Sprintf("redis eval error: %v", err))
	}
	n, ok := res.(int64)
	if !ok {
		return 0, errors.E(op, errors.Internal, fmt.Sprintf("unexpected return value from Redis script: %v", res))
	}
	return n, nil
}
