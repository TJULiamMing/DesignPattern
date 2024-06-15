package main

import "designPattern/BehavioralPattern/observer"

func main() {
	observer1 := observer.Customer{Id: "1"}
	observer2 := observer.Customer{Id: "2"}

	iphone := observer.NewItem("iphone")
	iphone.Register(&observer1)
	iphone.Register(&observer2)
	iphone.NotifyAll()

	iphone.Deregister(&observer2)
	iphone.UpdateAvailability()
}
