package reuse_test

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/valteem/reuse"
)

func TestUnmarshalRawJSON(t *testing.T) {
	b := []byte(`{"flightid":"QZ556", "origin":"DPS", "destination":"KUL", "flight_time":"4hrs"}`)
	err := reuse.UnmarshalRawJSON(&b)
	assert.Equal(t, nil, err, "should be equal")

	flightData := []byte(`
	{
    "aircraft":{"model":"A320-200", "registration":"QZ-ABC", "age":5},
	"flightid":"QZ556",
	"origin":"DPS",
	"destination":"KUL",
	"flight_time":"4hrs"
     }
	`)
	a, e := reuse.ReadAircraftFeatures(&flightData)
	if e != nil {
		fmt.Println(e)
	}
	assert.Equal(t, "A320-200", a.Model, "should be equal")
	assert.Equal(t, "QZ-ABC", a.Registration, "should be equal")
	assert.Equal(t, 5, a.Age, "should be equal")
}