package interview

import (
	"fmt"
	"runtime"
	"sync"
)

func Routine3() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("A: ", i)
			wg.Done()

		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()

		}(i)
	}
	wg.Wait()
}

// select 会随机选出一个case 执行
// 注意 chan

func Routine5() {
	runtime.GOMAXPROCS(1)
	intChan := make(chan int, 0)
	strChan := make(chan string, 0)
	intChan <- 1
	strChan <- "2"
	select {
	case value := <-intChan:
		fmt.Println(value)
	case <-strChan:
		panic("2")
	}

}
