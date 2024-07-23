package future

import (
	"context"
	"time"

	"github.com/go-errors/errors"
)

type IFuture[T any] interface {
	Get() T
	Set(T)
	getChan() chan T
}

type future[T any] struct {
	result chan T
}

func (f *future[T]) Get() T {
	return <-f.result
}

func (f *future[T]) getChan() chan T {
	return f.result
}

func (f *future[T]) Set(result T) {
	f.result <- result
}

func New[T any](t ...T) IFuture[T] {
	f := &future[T]{
		result: make(chan T, 1),
	}
	if len(t) > 0 {
		f.Set(t[0])
	}
	return f
}

// 等待很多个Future中一个的结果（error）。
// 无法等待slice，因此通过default自旋
func WaitOneError(ctx context.Context, futures ...IFuture[error]) (err error) {
	defer func() {
		if err != nil {
			err = errors.New(err)
		}
	}()

	for {
		for _, f := range futures {
			select {
			case err = <-f.getChan():
				return
			case <-ctx.Done():
				err = ctx.Err()
				return
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}
