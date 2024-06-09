package flyweight

type Dress interface {
	GetColor() string
}

type TerroristDress struct {
	color string
}

func (t *TerroristDress) GetColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}

type CounterTerroristDress struct {
	color string
}

func (c *CounterTerroristDress) GetColor() string {
	return c.color
}

func newCounterTerroristDress() *CounterTerroristDress {
	return &CounterTerroristDress{color: "green"}
}
