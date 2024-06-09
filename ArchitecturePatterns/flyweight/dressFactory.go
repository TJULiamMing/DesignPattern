package flyweight

import "fmt"

const (
	//TerroristDressType terrorist dress type
	TerroristDressType = "tDress"
	//CounterTerrroristDressType terrorist dress type
	CounterTerrroristDressType = "ctDress"
)

var (
	dressFactorySingleInstance = &DressFactory{
		DressMap: make(map[string]Dress),
	}
)

type DressFactory struct {
	DressMap map[string]Dress
}

func (d *DressFactory) getDressByType(dressType string) (Dress, error) {
	if d.DressMap[dressType] != nil {
		return d.DressMap[dressType], nil
	}

	if dressType == TerroristDressType {
		d.DressMap[dressType] = newTerroristDress()
		return d.DressMap[dressType], nil
	}
	if dressType == CounterTerrroristDressType {
		d.DressMap[dressType] = newCounterTerroristDress()
		return d.DressMap[dressType], nil
	}

	return nil, fmt.Errorf("Wrong dress type passed")
}

func GetDressFactorySingleInstance() *DressFactory {
	return dressFactorySingleInstance
}
