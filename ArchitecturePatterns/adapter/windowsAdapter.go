package adapter

import "fmt"

type WindowsAdapter struct {
	Windows *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.Windows.InsertIntoUSBPort()
}
