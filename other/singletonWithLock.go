package other

import "sync"

type singletonLock struct {
}

var insLock *singletonLock
var mu sync.Mutex

func GetInsWithLock() *singletonLock {
	mu.Lock()
	defer mu.Unlock()
	if insLock == nil {
		insLock = &singletonLock{}
	}
	return insLock
}

// lock 优化
func GetdoubleSingletionLock() *singletonLock {

	if insLock == nil {
		mu.Lock()
		defer mu.Unlock()
		insLock = &singletonLock{}
	}
	return insLock
}

// sync.once
var once sync.Once

func GetInsSync() *singletonLock {
	once.Do(func() {
		insLock = &singletonLock{}
	})
	return insLock
}
