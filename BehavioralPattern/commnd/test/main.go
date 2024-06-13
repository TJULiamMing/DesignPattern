package main

import "designPattern/BehavioralPattern/commnd"

func main() {
	light := commnd.Light{}
	rc := &commnd.RemoteControl{}

	OnCommand := commnd.LightOn{L: light}
	OffCommand := commnd.LightOff{L: light}

	rc.SetCommand(&OnCommand)
	rc.Press()

	rc.SetCommand(&OffCommand)
	rc.Press()

}
