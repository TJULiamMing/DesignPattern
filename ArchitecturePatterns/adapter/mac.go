package adapter

import "fmt"

type Mac struct{}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Successfully Insert Into Lightning Port")
}
