package reuse_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	reuse "github.com/valteem/reuse"
)

func TestConvertAssert(t *testing.T) {
	p := reuse.Person{"name", 30}
	i := reuse.ConvertPersonToIface(&p)
	a := reuse.ConvertIfaceToPerson(i)
	assert.Equal(t, "name", a.Name, "should be equal")
	assert.Equal(t, 30, a.Age, "should be equal")
} 	