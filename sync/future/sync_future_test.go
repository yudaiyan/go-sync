package future

import (
	"context"
	"testing"
	"time"

	"github.com/go-errors/errors"
)

func runJobA() IFuture[error] {
	f := New[error]()

	go func() {
		time.Sleep(1 * time.Second)
		f.Set(errors.New("超时A"))
	}()

	return f
}

func runJobB() IFuture[error] {
	f := New[error]()

	go func() {
		time.Sleep(1 * time.Second)
		f.Set(errors.New("超时B"))
	}()

	return f
}

func TestFuture(t *testing.T) {
	future := runJobA()
	// 其他 job
	if err := future.Get(); err != nil {
		t.Log(err.(*errors.Error).ErrorStack())
	}
}

func TestWaitError(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()
	err := WaitOneError(ctx, runJobA(), runJobB())
	if err != nil {
		t.Log(err.(*errors.Error).ErrorStack())
	}
}
