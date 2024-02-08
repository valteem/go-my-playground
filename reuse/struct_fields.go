// https://stackoverflow.com/a/58119914

package reuse

import (
	"time"
)

type Item struct {
	Name string
}

type ItemWrapper struct {
	Item      // promoted field
	time.Time // anonymous field
}