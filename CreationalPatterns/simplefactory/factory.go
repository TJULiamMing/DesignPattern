package simplefactory

// 工厂接口
type Factory interface {
	MakeProduct(name string) Product
}
