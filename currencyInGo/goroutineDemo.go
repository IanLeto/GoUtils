package currencyInGo

import (
	"fmt"
	"time"
)

func DoWork(strChan <-chan string) <-chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("work done")
		defer close(completed)
		for value := range strChan {
			fmt.Println(value)
		}
	}()
	return completed
}

// 这里strchan 是否只读是无所谓的
func DoWorkWithNoLeak(strChan chan string, done chan interface{}) chan interface{} {
	completed := make(chan interface{})
	go func() {
		defer fmt.Println("work done")
		defer close(completed)
		for {
			select {
			case s := <-strChan:
				fmt.Println(s)
			case <-time.After(1 * time.Second):
				fmt.Println("time out") //超时

			case <-done:
				return // 关闭goroutine 的信号
			}
		}
	}()
	return completed
}
