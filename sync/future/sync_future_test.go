package future

import (
	"errors"
	"testing"
	"time"
)

func runJobA() IFuture[error] {
	f := New[error]()

	go func() {
		time.Sleep(1 * time.Second)
		f.Set(errors.New("超时"))
	}()

	return f
}

func TestFuture(t *testing.T) {
	t.Log("running A")
	future := runJobA()
	t.Log("running Other")
	// 其他 job
	if err := future.Get(); err != nil {
		t.Log(err)
	}
}
