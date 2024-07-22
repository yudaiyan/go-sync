package sync

import (
	"net"
	"sync"
	"testing"
	"time"
)

func TestSyncMap(t *testing.T) {
	var m ISyncMap[string, int] = &SyncMap[string, int]{sync.Map{}}
	cready := make(chan bool)

	var print = func() {
		t.Log("-----")
		for k, v := range m.Iter() {
			t.Log(k, v)
		}
	}

	go func() {
		for i := 0; i < 2; i++ {
			<-cready
			print()
			cready <- true
		}
	}()

	m.Store("a", 123)
	cready <- true
	<-cready

	m.Store("e", 123)
	cready <- true
	<-cready

	time.Sleep(1 * time.Second)
}

func TestSyncMapComparableVal(t *testing.T) {
	tests := []struct {
		name string
	}{
		// 初始化第一个测试用例
		{
			name: "Test1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var m ISyncMapComparableVal[string, net.HardwareAddr, [6]byte] = &SyncMapComparableVal[string, net.HardwareAddr, [6]byte]{
				ToComparable: func(in net.HardwareAddr) [6]byte {
					var out [6]byte
					copy(out[:], in)
					return out
				},
			}
			m.Store("a", net.HardwareAddr{0x11, 0x22, 0x33, 0x44, 0x55, 0x66})
			m.Store("b", net.HardwareAddr{0x22, 0x22, 0x33, 0x44, 0x55, 0x66})

			t.Log(m.ContainsVal(net.HardwareAddr{0x22, 0x22, 0x33, 0x44, 0x55, 0x66}))
			t.Log(m.ContainsVal(net.HardwareAddr{0x33, 0x22, 0x33, 0x44, 0x55, 0x66}))
		})
	}
}
