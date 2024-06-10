package main

import "designPattern/BehavioralPattern/strategy"

func main() {
	creditCardStrategy := &strategy.CreditCardStrategy{
		Name:       "lpm",
		CardNumber: "1234 5678 9012 3456",
		Cvv:        "123",
		Date:       "01/22",
	}
	paymentContext := strategy.NewExecutePay(100, creditCardStrategy)
	if err := paymentContext.Pay(); err != nil {
		return
	}

	payPalStrategy := &strategy.PayPalStrategy{
		Email:    "123456789@qq.com",
		Password: "my.password",
	}
	paymentContext = strategy.NewExecutePay(200, payPalStrategy)

	if err := paymentContext.Pay(); err != nil {
		return
	}

}
