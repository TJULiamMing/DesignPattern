package main

import (
	"designPattern/CreationalPatterns/builder"
	"fmt"
)

// 生成器模式：具体构造过程由director封装实现
// 下面封装了villa的构建过程
// 如需拓展多个产品，只要实现相应builder接口（例如apartment.go）
func main() {
	villabuilder := builder.Villabuilder{}

	director := builder.NewDirector(&villabuilder)
	director.BuildHouse()

	fmt.Println(villabuilder.GetHouse())
}
