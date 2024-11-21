package main

import "sync"

type DB struct {
	count int
	mu    *sync.Mutex
}

func (db *DB) Query() int {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.count++
	return db.count
}
