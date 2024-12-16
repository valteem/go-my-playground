package binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Header struct {
	StoreID *int `header:"storeid" binding:"required"`
}

type Store struct {
	Location string `json:"location" binding:"required"`
	Square   int    `json:"square" binding:"required"`
}

func HandleNewStore(c *gin.Context) {

	store := &Store{}
	h := &Header{}

	if err := c.ShouldBindHeader(h); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := c.ShouldBindJSON(store); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, store)

}
