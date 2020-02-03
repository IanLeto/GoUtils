package utils

import (
	"fmt"
	"sync"
)

// 如果我们对普通map 进行读写操作，回发生竞争状态
func DemoSyncMap() {
	m := make(map[string]interface{})
	go func() {
		for {
			m["k"] = "v"
		}
	}()

	go func() {
		for {
			_ = m["k"]
		}
	}()

}

// 用sync map 来替代这种情况
func DemoSyncMap2() {
	var syncDict sync.Map
	syncDict.Store("k1", "v1")
	v, ok := syncDict.Load("k1")
	if ok {
		fmt.Println(v)
	}
	syncDict.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
