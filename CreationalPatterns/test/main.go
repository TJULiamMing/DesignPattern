package main

import (
	"designPattern/CreationalPatterns/Singleton"
	"fmt"
)

func main() {
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

	_, err := fmt.Scanln()
	if err != nil {
		return
	}

}
