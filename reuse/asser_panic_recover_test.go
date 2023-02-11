package reuse_test

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/valteem/reuse"
)

func TestCheckTypePerson(t *testing.T) {
	p := &reuse.Person{"somename", 30}
	var i interface{} = p
	c, err := reuse.CheckTypePerson(i)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, p.Name, c.Name, "should be equal")
	assert.Equal(t, p.Age, c.Age, "should be equal")

	w := struct{
		input int
		output int
	}{30, 30}
	i = w
	r, e := reuse.CheckTypePerson(i)
	if e != nil {
		fmt.Println(r, e)
	}
	assert.Equal(t, "", r.Name, "should be equal")
	assert.Equal(t, 0, r.Age, "should be equal")

}