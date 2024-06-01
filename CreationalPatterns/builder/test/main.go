package main

import (
	"designPattern/CreationalPatterns/builder"
	"fmt"
)

// 生成器模式：具体构造过程由director封装实现
// 下面封装了villa的构建过程
// 如需拓展多个产品，只要实现相应builder接口（例如apartment.go）
func main() {
	//获取指定产品生成器
	villabuilder := builder.Villabuilder{}

	//由director封装构造过程
	director := builder.NewDirector(&villabuilder)
	director.BuildHouse()

	//返回构造结果
	fmt.Println(villabuilder.GetHouse())
}
