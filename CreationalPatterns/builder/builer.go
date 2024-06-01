package builder

// 抽象生成器接口（例如：房子）
type Builder interface {
	BuildWall()
	BuildDoor()
	BuildWindow()
	GetHouse() string
}
