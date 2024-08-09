package server

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
)

type item struct {
	ID   int    `json:"id"`
	Desc string `json:"desc"`
}

var (
	items = map[int]*item{} // Id is both map key and struct field
	seq   = 1               // numbering
	lock  = sync.Mutex{}
)

func CreateItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	it := &item{
		ID: seq,
	}
	if err := c.Bind(it); err != nil {
		return err
	}
	items[it.ID] = it
	seq++
	return c.JSON(http.StatusCreated, it)
}
