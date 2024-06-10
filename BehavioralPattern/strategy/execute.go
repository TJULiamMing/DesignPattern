package strategy

type ExecutePay struct {
	amount   float64
	strategy PayStrategy
}

func NewExecutePay(amount float64, strategy PayStrategy) *ExecutePay {
	return &ExecutePay{
		amount:   amount,
		strategy: strategy,
	}
}

func (p *ExecutePay) Pay() error {
	return p.strategy.Pay(p.amount)
}
