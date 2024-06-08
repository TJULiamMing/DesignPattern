package main

import "designPattern/ArchitecturePatterns/facade"

func main() {
	cf := facade.NewComputerFacaxde(2.8, 16, 1024)
	cf.Run()
}
