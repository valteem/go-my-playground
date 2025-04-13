package sqldb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	defaultMaxPoolSize  = 1
	defaultConnAttempts = 10
	defaultConnTimeout  = time.Second
)

type PostgresDB struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Pool *pgxpool.Pool
}

func New(connString string, opts ...Option) (*PostgresDB, error) {

	pg := &PostgresDB{
		maxPoolSize:  defaultMaxPoolSize,
		connAttempts: defaultConnAttempts,
		connTimeout:  defaultConnTimeout,
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			opt(pg)
		}
	}

	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		// error formatting (%s vs %v vs %w):
		// https://stackoverflow.com/a/61287626
		return nil, fmt.Errorf("failed to parse pool config: %w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for pg.connAttempts > 0 {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		pg.connAttempts--
		log.Printf("trying to connect to database, %d attempts left", pg.connAttempts)
		time.Sleep(pg.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to establish database connection: %w", err)
	}

	return pg, nil

}

func (p *PostgresDB) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

type Option func(*PostgresDB)

func MaxPoolSize(size int) Option {
	return func(p *PostgresDB) {
		p.maxPoolSize = size
	}
}

func ConnAttempts(attempts int) Option {
	return func(p *PostgresDB) {
		p.connAttempts = attempts
	}
}

func ConnTimeout(t time.Duration) Option {
	return func(p *PostgresDB) {
		p.connTimeout = t
	}
}
