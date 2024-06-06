package decorator

type ExtraDecorator struct {
	ExtraPrice float32
}

func (e ExtraDecorator) DecoratePrice(car Car) Car {
	car.Price += e.ExtraPrice
	return car
}
