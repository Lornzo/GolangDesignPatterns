package example3

type IStyle interface {
	StyleDescription() string
}

type CowboyStyle struct{}

func (c CowboyStyle) StyleDescription() string {
	return "Cowboy Style"
}

type TaiwanStyle struct{}

func (t TaiwanStyle) StyleDescription() string {
	return "Taiwan Style"
}
