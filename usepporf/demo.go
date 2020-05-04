package usepporf

import (
	"GoUtils/utils"
	"fmt"
	"github.com/pkg/profile"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func Demopprof() {

}

func WorkerManager(n int) {
	fmt.Println("start")
	for i := 0; i < n; i++ {
		go worker()

	}
	time.Sleep(20 * time.Second)
	utils.NoErr(http.ListenAndServe("0.0.0.0:6060", nil))

}

func worker() {
	time.Sleep(1 * time.Second)
}

func DemoSlice() {
	defer profile.Start().Stop()

}

func makeSlice() []int {
	sl := make([]int, 100000000)
	for index := range sl {
		sl[index] = index
	}
	return sl
}

func sumSlice(sl []int) int {
	sum := 0
	for _, x := range sl {
		sum += x
	}
	return sum
}
func init() {
}
