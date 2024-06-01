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

/*
简单工厂：无工厂接口，实现一个具体工厂，一个产品接口，多个具体产品	->	缺点：增加产品类会修改具体工厂的代码
工厂方法：一个工厂接口，实现多个具体工厂，一个产品接口，多个具体产品
抽象工厂：一个工厂接口，实现多个具体工厂，多个产品接口，多个具体产品
*/
