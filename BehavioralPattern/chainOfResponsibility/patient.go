package chainOfResponsibility

type Patient struct {
	Name             string
	RegistrationDone bool
	ClinicCheckDone  bool
	PaymentDone      bool
}

type department interface {
	Operate(Patient)
	SetNext(department)
}
