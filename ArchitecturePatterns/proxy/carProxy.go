package proxy

import "fmt"

type Driver struct {
	Age int
}

type CarProxy struct {
	car    Car
	driver Driver
}

func NewCarProxy(driver Driver) *CarProxy {
	return &CarProxy{driver: driver}
}

func (c *CarProxy) Drive() {
	if c.driver.Age >= 18 {
		c.car.Run()
	} else {
		fmt.Println("driver is too young to drive")
	}
}
