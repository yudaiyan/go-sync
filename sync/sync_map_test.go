package sync

import (
	"log"
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
				log.Println(k, v)
			}
			log.Println("-----")

			m.Store("e", 567)
			for k, v := range m.Iter() {
				log.Println(k, v)
			}

			log.Println("-----")

			m.Delete("e")
			for k, v := range m.Iter() {
				log.Println(k, v)
			}

			log.Println("-----")
			log.Print(m.Load("a"))
		})
	}
}
