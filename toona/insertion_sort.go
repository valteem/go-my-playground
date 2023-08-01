package toona

import (
	"fmt"
)

// comp(t1, t2) returns true if t1 < t2
func InsertionSort [T any] (s []T, comp func(t1, t2 T) bool) error {

	length := len(s)
	if length < 2 {
		return fmt.Errorf("slice length must be greater than 1")
	}

	for j := 1; j < length; j++ {
		key := s[j]
		i := j - 1
		for {
			if !( (i > -1) && comp(key, s[i]) ) {
				break
			}
			s[i+1] = s[i]
			i--
		}
		s[i+1] = key
	} 

	return nil

}