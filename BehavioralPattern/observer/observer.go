package observer

import "fmt"

type Observer interface {
	Update(string)
}

type Customer struct {
	Id string
}

func (c *Customer) Update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.Id, itemName)
}
