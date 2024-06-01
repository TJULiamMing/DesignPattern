package factorymethod

// 工厂接口
type Factory interface {
	MakeProduct() (Product, error)
}
