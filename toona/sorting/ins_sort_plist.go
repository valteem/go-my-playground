// insertion sort of positional list

package sorting

import(
	"reflect"
	"my.play.go/toona/sequences"
)

func InsSortPosList[T any](pl *sequences.PList[T], gt func(T, T) bool) {

	tail, _ := pl.Tail() // assuming silently that positional list has non zero length
	point := tail
	head, _ := pl.Head()
	for !reflect.DeepEqual(point, head) {
		handle, _ := pl.After(point)
		elt := handle.Element()
		if gt(elt, point.Element()) {
			point = handle
		} else {
			move := point
			for {
				before, _ := pl.Before(move)
				if !gt(before.Element(), elt) {
					break
				}
				move = before
				if move == tail {
					break
				}
			}
			pl.Remove(handle)
			pl.InsertBefore(elt, move)
		}
		head , _ = pl.Head()
	}
	
}