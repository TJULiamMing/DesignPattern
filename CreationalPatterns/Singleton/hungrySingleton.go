package Singleton

import "fmt"

// 饿汉式
// 导入包时立即创建对象
// 缺点：对象不使用时占用内存

type HungrySingleton struct {
	value int
}

var Singleton2 *HungrySingleton

func init() {
	if Singleton2 == nil {
		fmt.Println("new hungrySingleton")
		Singleton2 = new(HungrySingleton)
	}
}

func GethungrySingleton() *HungrySingleton {
	return Singleton2
}
