package chainOfResponsibility

import "fmt"

// 前台
type Clinic struct {
	Next department
}

func (c *Clinic) Operate(patient Patient) {
	if patient.RegistrationDone {
		fmt.Println("已就诊过，下一步：结账")
		c.Next.Operate(patient)
		return
	}

	fmt.Println("正在就诊")
	patient.RegistrationDone = true
	c.Next.Operate(patient)
	return
}

func (c *Clinic) SetNext(next department) {
	c.Next = next
}
