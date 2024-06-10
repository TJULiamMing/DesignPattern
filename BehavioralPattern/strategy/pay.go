package strategy

import "fmt"

// 支付接口
type PayStrategy interface {
	Pay(amount float64) error
}

// 实现具体的支付策略
type CreditCardStrategy struct {
	Name       string
	CardNumber string
	Cvv        string
	Date       string
}

func (c *CreditCardStrategy) Pay(amount float64) error {
	fmt.Printf("Paying %0.2f using credit card\n", amount)
	return nil
}

type PayPalStrategy struct {
	Email    string
	Password string
}

func (p *PayPalStrategy) Pay(amount float64) error {
	fmt.Printf("Paying %0.2f using PayPal\n", amount)
	return nil
}
