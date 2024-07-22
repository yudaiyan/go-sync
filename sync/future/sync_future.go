package future

type IFuture[T any] interface {
	Get() T
	Set(T)
}

type future[T any] struct {
	result chan T
}

func (f *future[T]) Get() T {
	return <-f.result
}

func (f *future[T]) Set(result T) {
	f.result <- result
}

func New[T any]() IFuture[T] {
	return &future[T]{
		result: make(chan T),
	}
}
