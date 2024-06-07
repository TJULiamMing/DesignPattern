package proxy

import "fmt"

type Car struct{}

type Vehicle interface {
	Run()
}

func (c *Car) Run() {
	fmt.Println("Car is running")
}
