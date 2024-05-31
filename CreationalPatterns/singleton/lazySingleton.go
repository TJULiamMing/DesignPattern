package singleton

import (
	"fmt"
	"sync"
)

// 懒汉式
// 创建对象时不直接创建对象，获取对象时才创建对象
// 缺点：存在并发问题
type LazySingleton struct {
	value int
}

var Singleton1 *LazySingleton

func GetlazySingleton() *LazySingleton {
	if Singleton1 == nil {
		fmt.Println("new lazySingleton")
		Singleton1 = new(LazySingleton)
	} else {
		fmt.Println("lazySingleton is already initialized")
	}
	return Singleton1
}

//进阶：通过mutex双重检验解决并发问题，并且性能不会明显降低

var mutex sync.Mutex

var Singleton4 *LazySingleton

func GetlazySingleton_mutex() *LazySingleton {
	if Singleton4 == nil {
		mutex.Lock()
		if Singleton4 == nil {
			fmt.Println("new lazySingleton_mutex")
			Singleton4 = new(LazySingleton)
		}
		mutex.Unlock()
	} else {
		fmt.Println("lazySingleton_mutex is already initialized")
	}
	return Singleton4
}
