package main

import "designPattern/BehavioralPattern/chainOfResponsibility"

func init() {

}

func main() {
	cashier := &chainOfResponsibility.Cashier{Next: nil}
	clinic := &chainOfResponsibility.Clinic{Next: cashier}
	reception := &chainOfResponsibility.Reception{Next: clinic}

	patient := chainOfResponsibility.Patient{
		Name:             "lpm",
		RegistrationDone: false,
		ClinicCheckDone:  false,
		PaymentDone:      false,
	}

	reception.Operate(patient)
}
