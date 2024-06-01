package factorymethod

// 工厂接口
// 工厂方法模式：一个抽象产品类，可以派生出多个具体产品类。每个具体工厂类只能创建一个具体产品类的实例。
type Factory interface {
	MakeProduct() (Product, error)
}
