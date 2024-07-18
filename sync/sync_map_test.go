package sync

import (
	"net"
	"testing"
)

func TestSyncMap(t *testing.T) {
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
			var m ISyncMap[string, int] = &SyncMap[string, int]{}
			m.Store("a", 123)
			m.Store("b", 234)
			m.Store("c", 345)
			m.Store("d", 456)
			for k, v := range m.Iter() {
				t.Log(k, v)
			}
			t.Log("-----")

			m.Store("e", 567)
			for k, v := range m.Iter() {
				t.Log(k, v)
			}

			t.Log("-----")

			m.Delete("e")
			for k, v := range m.Iter() {
				t.Log(k, v)
			}

			t.Log("-----")
			t.Log(m.Load("a"))
		})
	}
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
