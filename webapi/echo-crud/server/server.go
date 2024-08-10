package server

import (
	"net/http"
	"strconv"
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

func GetItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	it, ok := items[id]
	if ok {
		return c.JSON(http.StatusOK, it)
	} else {
		return c.JSON(http.StatusNotFound, nil)
	}
}

func UpdateItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	it, ok := items[id]
	if ok {
		if err := c.Bind(it); err != nil {
			return err
		}
		items[id] = it
		return c.JSON(http.StatusOK, it)
	}
	return c.JSON(http.StatusNotFound, nil)
}

func DeleteItem(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	id, _ := strconv.Atoi(c.Param("id"))
	_, ok := items[id]
	if ok {
		delete(items, id)
		return c.NoContent(http.StatusNoContent)
	}
	return c.JSON(http.StatusNotFound, nil)
}
