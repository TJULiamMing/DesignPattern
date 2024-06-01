package factorymethod

// 具体工厂：soccerfactory
type SoccerFactory struct {
}

func (sf *SoccerFactory) MakeProduct() (Product, error) {
	return &Basketball{name: "soccer"}, nil
}
