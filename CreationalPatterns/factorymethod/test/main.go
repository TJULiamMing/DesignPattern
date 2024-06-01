package main

import (
	"designPattern/CreationalPatterns/factorymethod"
	"fmt"
)

// 工厂方法
func main() {

	bf := &factorymethod.BasketballFactory{}
	basetball, _ := bf.MakeProduct()
	fmt.Println(basetball.Use())

	sf := &factorymethod.SoccerFactory{}
	soccer, _ := sf.MakeProduct()
	fmt.Println(soccer.Use())

}
