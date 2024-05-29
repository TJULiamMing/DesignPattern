package main

import (
	"fmt"
	"sigleton1"
)

func main() {
	for i:=0;i<10;i++ {
		go sigleton1.GetmySingleton()
	}

	fmt.Scanln()
}