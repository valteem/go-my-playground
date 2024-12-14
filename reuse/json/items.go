package json

import (
	"fmt"
)

type ItemStorage interface {
	Get(string) (*Item, bool)
}

// In-memory storage: implements ItemStorage interface
type memItemStorage map[string]*Item

func (m memItemStorage) Get(id string) (*Item, bool) {
	item, ok := m[id]
	return item, ok
}

// Create new empty in-memory storage
func NewMemItemSorage() memItemStorage {
	return memItemStorage{}
}

// Initialize memory item storage with arbitrary data
func (m memItemStorage) init(size int) {
	for i := 1; i <= size; i++ {
		id := fmt.Sprintf("i%03d", i)
		desc := fmt.Sprintf("item #%04d", i)
		m[id] = &Item{Name: id, Description: desc}
	}
}
