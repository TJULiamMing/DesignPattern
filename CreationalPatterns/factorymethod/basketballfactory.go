package factorymethod

// 具体工厂：basketballfactory
type BasketballFactory struct {
}

func (bf *BasketballFactory) MakeProduct() (Product, error) {
	return &Basketball{name: "basketball"}, nil
}
