package factorymethod

// 产品接口
// 工厂方法只有一个抽象产品类
// 补充:抽象工厂模式有多个抽象产品类，每个抽象产品类可以派生出多个具体产品类。一个抽象工厂类可以派生出多个具体工厂类。每个具体工厂类可以创建多个具体产品的实例。
type Product interface {
	Use() string
}
