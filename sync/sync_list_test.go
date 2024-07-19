package sync

import (
	"testing"
)

func TestA(t *testing.T) {
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
			var list ISyncList[string] = &SyncList[string]{}
			list.Add("a")
			list.Add("b")
			list.Add("c")
			for _, item := range list.Iter() {
				t.Log(item)
			}
			t.Log("-----")

			list.Add("d")
			for _, item := range list.Iter() {
				t.Log(item)
			}

			t.Log("-----")

			list.Remove(0)
			for _, item := range list.Iter() {
				t.Log(item)
			}

			t.Log("-----")
			t.Log(list.Get(0))

			t.Log("-----")
			t.Log(list.Size())

			list.RemoveByVal("b")
			t.Log(list)

		})
	}
}
