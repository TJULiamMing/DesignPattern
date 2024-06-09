package flyweight

type Player struct {
	dress      Dress
	playerType string
	lat        int
	long       int
}

func NewPlayer(playerType, dressType string) *Player {
	dress, _ := GetDressFactorySingleInstance().getDressByType(dressType)
	return &Player{
		playerType: playerType,
		dress:      dress,
	}
}

func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
