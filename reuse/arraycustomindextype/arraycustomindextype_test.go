package arraycustomindextype

import (
	"testing"
)

func TestArrInit(t *testing.T) {

	for i := range 5 {
		if squares[i] != i*i {
			t.Errorf("#%d square: get %d, expect %d", i, squares[i], i*i)
		}
	}

}
