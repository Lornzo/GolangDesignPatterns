package example3

type IExchangeRate interface {
	ExchangeRate() float64
}

type USDExchangeRate struct{}

func (u USDExchangeRate) ExchangeRate() float64 {
	return 1
}

type TWExchangeRate struct{}

func (t TWExchangeRate) ExchangeRate() float64 {
	return 30.5
}
