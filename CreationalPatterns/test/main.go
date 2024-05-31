package main

import (
	"designPattern/CreationalPatterns/Singleton"
	"fmt"
)

func singleton_test() {
	for i := 0; i < 5; i++ {
		go Singleton.GetlazySingleton()
	}

	for i := 0; i < 5; i++ {
		go Singleton.GetlazySingleton_mutex()
	}

	for i := 0; i < 5; i++ {
		go Singleton.GethungrySingleton()
	}

	for i := 0; i < 5; i++ {
		go Singleton.GetonceSingleton()
	}

	return

}
func main() {
	singleton_test()

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
