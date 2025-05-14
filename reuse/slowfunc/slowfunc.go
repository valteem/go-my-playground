package slowfunc

import (
	"context"
)

func WrapSlowFunc[T any](ctx context.Context, slowfunc func(args ...any) T, args ...any) (T, error) {

	ch := make(chan T, 1)
	go func() {
		ch <- slowfunc(args...)
	}()

	select {
	case v := <-ch:
		return v, nil
	case <-ctx.Done():
		var v T
		return v, ctx.Err()
	}

}
