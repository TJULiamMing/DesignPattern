package chainOfResponsibility

import "fmt"

// 前台
type Cashier struct {
	Next department
}

func (c *Cashier) Operate(patient Patient) {
	if patient.RegistrationDone {
		fmt.Println("已结账，无下一步")
		return
	}

	fmt.Println("正在结账，结账完成后离开")
	patient.RegistrationDone = true
	return
}

func (c *Cashier) SetNext(next department) {
	c.Next = next
}
