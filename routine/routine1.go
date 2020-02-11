package routine

import "sync"

type xx struct {
}

var test *xx
var mu sync.Mutex

func adiaojf() *xx {
	mu.Lock()
	defer mu.Unlock()
	if test == nil {
		test = &xx{}
	}
	return test

}

var once sync.Once

func afeioawe() *xx {
	once.Do(func() {
		test = &xx{}
	})
	return test
}
