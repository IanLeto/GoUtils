package other

type singleton struct {
}

// 单例模式：单例对象必须只有一个实例存在
// 懒汉: 第一次使用构建
var ins *singleton

// 非线程安全
func GetIns() *singleton {
	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

type hungrySingleton struct {
}

// 饿汉：装载时构建
var insHungry *hungrySingleton = &hungrySingleton{}

func GetnInsHungry() *hungrySingleton {
	return insHungry
}
