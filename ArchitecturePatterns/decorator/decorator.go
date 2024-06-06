package decorator

type Decorator interface {
	DecoratePrice(car Car) Car
}
