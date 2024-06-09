package main

import (
	"designPattern/ArchitecturePatterns/flyweight"
	"fmt"
)

func main() {
	game := flyweight.NewGame()

	//Add Terrorist
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)
	game.AddTerrorist(flyweight.TerroristDressType)

	//Add CounterTerrorist
	game.AddCounterTerrorist(flyweight.CounterTerrroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerrroristDressType)
	game.AddCounterTerrorist(flyweight.CounterTerrroristDressType)

	dressFactoryInstance := flyweight.GetDressFactorySingleInstance()

	for dressType, dress := range dressFactoryInstance.DressMap {
		fmt.Printf("DressColorType: %s\nDressColor: %s\n", dressType, dress.GetColor())
	}
}
