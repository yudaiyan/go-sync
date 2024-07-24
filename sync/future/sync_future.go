package future

import (
	"context"
	"slices"
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

// 等待很多个Future的结果：
//
// 1. 如果存在error，则立即返回一个error，不等待其他future执行完。
//
// 2. 如果不存在error，则需要等到所有future执行完成。
func WaitResult(ctx context.Context, futures ...IFuture[error]) (err error) {
	defer func() {
		if err != nil {
			err = errors.New(err)
		}
	}()

	for {
		for i, f := range futures {
			select {
			case err = <-f.getChan():
				if err != nil {
					return
				}
				// 排除已经完成且没有error的。
				futures = slices.Delete(futures, i, i+1)
				// 由于修改了futures，因此需要重新进入循环。
				goto outLoop
			case <-ctx.Done():
				err = ctx.Err()
				return
			default: // 无法在一个select中case slice的每个元素，因此通过default立即select下个元素
			}
		}
	outLoop:
		if len(futures) == 0 { // 表示所有future都有结果，且不存在nil
			return
		}
		time.Sleep(50 * time.Millisecond)
	}
}
