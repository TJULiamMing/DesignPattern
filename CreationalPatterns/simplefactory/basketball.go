package simplefactory

// 具体产品：basketball
type Basketball struct {
	name string
}

func (basketball *Basketball) Use() string {
	return "playing " + basketball.name
}
