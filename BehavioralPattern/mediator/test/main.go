package main

import (
	"designPattern/BehavioralPattern/mediator"
)

func main() {
	trainMediator := mediator.NewTrainMediator()

	p := &mediator.PassengerTrain{Mediator: trainMediator}
	f := &mediator.FreightTrain{Mediator: trainMediator}

	p.Arrive()
	f.Arrive()
	p.Depart()
}
