package main

import (
	"designPattern/ArchitecturePatterns/adapter"
	"fmt"
)

func main() {
	client := &adapter.Client{}
	mac := &adapter.Mac{}
	client.InsertLightningConnectorIntoComputer(mac)
	fmt.Println("下面使用适配器连接Windows")
	windows := &adapter.Windows{}
	wa := &adapter.WindowsAdapter{Windows: windows}
	client.InsertLightningConnectorIntoComputer(wa)
}
