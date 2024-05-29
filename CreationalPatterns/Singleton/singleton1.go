package pkg

import (
	"fmt"
)

// 懒汉式
// 创建对象时不直接创建对象，获取对象时才创建对象
// 缺点：存在并发问题
type singleton1 {
	value int
}

var mySingleton *singleton1 

func GetmySingleton() *singleton1 {
	if mySingleton == nil {
		fmt.Println("new mySingleton")
		mySingleton = new(singleton1)
	}
	return mySingleton
}