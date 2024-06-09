package flyweight

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 1),
		counterTerrorists: make([]*Player, 1),
	}
}

func (c *Game) AddTerrorist(dressType string) {
	player := NewPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *Game) AddCounterTerrorist(dressType string) {
	player := NewPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}
