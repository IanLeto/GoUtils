package utils

import (
	"fmt"
	"math/rand"
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
	cc, ok := syncDict.LoadOrStore("k1", "v2")
	if ok {
		fmt.Println("ok", cc)
	} else {
		fmt.Println(" k1 exist already", cc)
	}
}

func DemoSyncMap3() {
	var x sync.Map
	var count int
	loadStore := func() {
		n := rand.Intn(100000)
		x.Store(n, n+1)
	}
	VisitAll(loadStore)
	x.Range(func(key, value interface{}) bool {
		count++
		fmt.Println(key, value)
		return true
	})
	fmt.Println("total",count)

}
func VisitAll(fn func()) {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			fn()
		}(&wg)
	}
	wg.Wait()
}
