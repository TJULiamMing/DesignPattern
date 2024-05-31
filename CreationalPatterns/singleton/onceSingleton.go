package singleton

import (
	"fmt"
	"sync"
)

type OnceSingleton struct {
	value int
}

var singleton3 *OnceSingleton

var once sync.Once

func GetonceSingleton() *OnceSingleton {
	once.Do(func() {
		fmt.Println("new onceSingleton")
		singleton3 = new(OnceSingleton)
	})

	return singleton3
}
