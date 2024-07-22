package sync

import (
	"sync"
	"testing"
	"time"
)

// 验证 sync.Map 的可见性
func TestSyncMap0(t *testing.T) {
	m := sync.Map{}
	cready := make(chan bool)

	var print = func() func(any, any) bool {
		t.Log("--------")
		return func(key, value any) bool {
			t.Log(key, value)
			return true
		}
	}

	go func() {
		for i := 0; i < 2; i++ {
			<-cready
			m.Range(print())
			cready <- true
		}
	}()

	m.Store("a", 123)
	cready <- true
	<-cready

	m.Store("b", 123)
	cready <- true
	<-cready

	time.Sleep(1 * time.Second)
}
