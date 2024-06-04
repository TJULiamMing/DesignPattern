package adapter

import "fmt"

// 默认该client只能对接Lightning接口
type Client struct{}

func (c *Client) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
