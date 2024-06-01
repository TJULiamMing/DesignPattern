package simplefactory

// 具体产品：soccer
type Soccer struct {
	name string
}

func (soccer *Soccer) Use() string {
	return "playing " + soccer.name
}
