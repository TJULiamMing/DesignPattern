package mediator

import "fmt"

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}

// 客运火车
type PassengerTrain struct {
	Mediator Mediator
}

func (p *PassengerTrain) Arrive() {
	if !p.Mediator.CanArrive(p) {
		fmt.Println("PassengerTrain train blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain train arrived")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain Leaving")
	p.Mediator.NotifyDeparture()
}

func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain permit arrival")
	p.Arrive()
}

// 货运火车
type FreightTrain struct {
	Mediator Mediator
}

func (f *FreightTrain) Arrive() {
	if !f.Mediator.CanArrive(f) {
		fmt.Println("FreightTrain train blocked, waiting")
		return
	}
	fmt.Println("FreightTrain train arrived")
}

func (f *FreightTrain) Depart() {
	fmt.Println("FreightTrain Leaving")
	f.Mediator.NotifyDeparture()
}

func (f *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain permit arrival")
	f.Arrive()
}
