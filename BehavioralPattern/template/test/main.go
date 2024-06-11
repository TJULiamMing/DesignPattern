package main

import "designPattern/BehavioralPattern/template"

func main() {
	loler := template.LoL{Name: "lol"}
	template.Play(&loler)
}
