package main

import (
	"designPattern/CreationalPatterns/factorymethod"
	"fmt"
)

// 工厂方法
func main() {

	f := &factorymethod.BallFactory{}

	basetball, _ := f.MakeProduct("basketball")
	fmt.Println(basetball.Use())

	soccer, _ := f.MakeProduct("soccer")
	fmt.Println(soccer.Use())

	tennis, e := f.MakeProduct("tennis")
	if e != nil {
		fmt.Println(e.Error())
	} else {
		fmt.Println(tennis.Use())
	}
}
