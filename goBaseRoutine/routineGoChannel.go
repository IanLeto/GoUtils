package goBaseRoutine

import (
	"fmt"
	"math/rand"
	"sync"
)
// 什么垃圾题目？？
func xx() {
	var wg sync.WaitGroup
	ch := make(chan int, 0)
	wg.Add(1)
	go func(ch chan int) {
		ch <- rand.Intn(5)
	}(ch)

	go func() {
		defer wg.Done()
		for range ch {
			fmt.Println(<-ch)
		}
	}()
}
