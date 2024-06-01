package simplefactory

import "errors"

// 具体工厂：ballfactory
// 简单工厂方法缺点：添加或减少产品时，会污染（修改）MakeProduct方法
type BallFactory struct {
}

func (f *BallFactory) MakeProduct(ball string) (Product, error) {
	if ball == "basketball" {
		return &Basketball{name: ball}, nil
	} else if ball == "soccer" {
		return &Soccer{name: ball}, nil
	}

	return nil, errors.New("Invalid Ball")
}
