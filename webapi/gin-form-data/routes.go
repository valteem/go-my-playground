package main

import (
	"github.com/gin-gonic/gin"
)

func GetPerson(c *gin.Context) {
	var b Person
	c.Bind(&b)
	c.JSON(200, gin.H{
		"name": b.Name,
		"age":  b.Age,
		//		"contactinfo": b.ContactInfo,
	})
}

/* func GetAddress(c *gin.Context) {
	var b Address
	c.Bind(&b)
	c.JSON(200, gin.H{
		"zipcode":     b.ZIPCode,
		"city":        b.City,
		"street":      b.Street,
		"housenumber": b.HouseNumber,
	})
} */
