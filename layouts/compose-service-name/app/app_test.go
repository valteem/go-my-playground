package app

import (
	"context"

	"testing"

	"github.com/jackc/pgx/v5"
)

func TestBasicQuery(t *testing.T) {

	conn, err := pgx.Connect(context.Background(), ConnStr)
	if err != nil {
		t.Fatalf("failed to establish database connection: %v", err)
	}
	defer conn.Close(context.Background())

	var result int
	err = conn.QueryRow(context.Background(), "select 1 + 2").Scan(&result)
	if err != nil {
		t.Fatalf("failed to fetch basic query result: %v", err)
	}

	if actual, expected := result, 3; actual != expected {
		t.Errorf("query result: get %d, expect %d", actual, expected)
	}

}

func TestThatAlwaysFails(t *testing.T) {

	if 1 != 2 {
		t.Fatalf("1 == 2 always fails")
	}

}
