package utils

import "time"

func MapCopy(a, b map[string]interface{}) {
	for key, value := range a {
		b[key] = value
	}
}

func MapCopyValue(a map[string]interface{}) map[string]interface{} {
	var b = make(map[string]interface{}, 0)
	MapCopy(a, b)
	return b
}

var limit int // 最大并发数量
var count int // 业务下主机数量

func SearchHost() {
	time.Sleep(1 * time.Second)
}
func SearchTopo() {
	time.Sleep(2 * time.Second)
}
func Merge(a, b interface{}) interface{} {
	return a
}
func SchedulerSearch() {

}
