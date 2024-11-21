package main

import (
	"net/http"
	"strconv"
	"sync"
)

type Handler struct {
	pool *sync.Pool
}

func NewHandler(db *DB) *Handler {
	h := &Handler{pool: &sync.Pool{New: func() any { return NewConn(db) }}}
	return h
}

func (h *Handler) GetConn() *Conn {
	c := h.pool.Get().(*Conn)
	return c
}

func (h *Handler) PutConn(c *Conn) {
	h.pool.Put(c)
}

func (h *Handler) GetCount(w http.ResponseWriter, r *http.Request) {

	c := h.GetConn()
	defer h.PutConn(c)

	count := c.Query()
	w.Write([]byte(strconv.Itoa(count)))

}
