package chainOfResponsibility

import "fmt"

// 前台
type Reception struct {
	Next department
}

func (r *Reception) Operate(patient Patient) {
	if patient.RegistrationDone {
		fmt.Println("前台已经登记过，下一步：就诊")
		r.Next.Operate(patient)
		return
	}

	fmt.Println("正在前台登记")
	patient.RegistrationDone = true
	r.Next.Operate(patient)
	return
}

func (r *Reception) SetNext(next department) {
	r.Next = next
}
