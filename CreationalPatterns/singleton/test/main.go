package main

import (
	"designPattern/CreationalPatterns/singleton"
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		go singleton.GetlazySingleton()
	}

	for i := 0; i < 5; i++ {
		go singleton.GetlazySingleton_mutex()
	}

	for i := 0; i < 5; i++ {
		go singleton.GethungrySingleton()
	}

	for i := 0; i < 5; i++ {
		go singleton.GetonceSingleton()
	}

	return

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
