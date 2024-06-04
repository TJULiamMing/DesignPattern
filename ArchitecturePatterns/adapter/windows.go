package adapter

import "fmt"

type Windows struct{}

func (w *Windows) InsertIntoUSBPort() {
	fmt.Println("Successfully Insert Into USB Port")
}
