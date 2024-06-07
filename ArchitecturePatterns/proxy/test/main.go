package main

import (
	"designPattern/ArchitecturePatterns/proxy"
)

func main() {
	car1 := proxy.NewCarProxy(proxy.Driver{Age: 15})
	car1.Drive()

	car2 := proxy.NewCarProxy(proxy.Driver{Age: 25})
	car2.Drive()

}
