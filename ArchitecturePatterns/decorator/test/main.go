package main

import (
	"designPattern/ArchitecturePatterns/decorator"
	"fmt"
)

func main() {
	car := decorator.Car{Brand: "su7", Price: 200000}
	priceDecorator := decorator.ExtraDecorator{ExtraPrice: 50000}

	decoratorCar := priceDecorator.DecoratePrice(car)

	fmt.Println(decoratorCar)
}
