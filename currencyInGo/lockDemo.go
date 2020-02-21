package currencyInGo

import (
	"fmt"
	"sync"
	"time"
	"transfer/logging"
)

func DemoWithoutLock() {
	// 没能锁住临界资源导致最终count 不一定是0
	logger := logging.GetStdLogger()
	var count int

	//var lock sync.Mutex // 互斥锁：一个从公资源只能呗一个线程or进程访问
	increment := func() {
		//lock.Lock()
		//defer lock.Unlock()
		count++
		logger.Infof("incrementing: %d", count)
	}
	decrement := func() {
		//lock.Lock()
		//defer lock.Unlock()
		count--
		logger.Infof("decrementing: %d", count)
	}
	var arithmetic sync.WaitGroup
	for i := 0; i <= 15; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i <= 15; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()

}
func DemoLock() {
	logger := logging.GetStdLogger()
	var count int
	var lock sync.Mutex // 互斥锁：一个从公资源只能呗一个线程or进程访问
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		logger.Infof("incrementing: %d", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		logger.Infof("decrementing: %d", count)
	}
	var arithmetic sync.WaitGroup
	for i := 0; i <= 15; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}
	for i := 0; i <= 15; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()

}

func Demo2() {
	var lock sync.Mutex
	var xx int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		lock.Lock()
		defer lock.Unlock()
		xx = 1
		fmt.Println(xx)
		time.Sleep(10 * time.Second)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		xx = 12
		fmt.Println(xx)
	}()
	wg.Wait()
}

func DemoCond() {

}
func DemoOnce() {
	var (
		count int
		once  sync.Once
		wg    sync.WaitGroup
	)

	increment := func() {
		count++
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// once 会统计do中的方法来确保其只调用一次
			once.Do(increment)
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

func DemoOnce2() {
	var (
		count int
		once  sync.Once
	)
	xx := func() {
		count++
	}
	once.Do(xx)
	once.Do(xx)
	// 统计Do中的方法
	fmt.Println(count)
}
