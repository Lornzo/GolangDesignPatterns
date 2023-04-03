package example3

type ICountry interface {
	CreateMenu() IMenu
	CreateExchangeRate() IExchangeRate
	CreateStyle() IStyle
}

type USA struct{}

func (u USA) CreateMenu() IMenu {
	return EnglishMenu{}
}

func (u USA) CreateExchangeRate() IExchangeRate {
	return USDExchangeRate{}
}

func (u USA) CreateStyle() IStyle {
	return CowboyStyle{}
}

type Taiwan struct{}

func (t Taiwan) CreateMenu() IMenu {
	return ChineseMenu{}
}

func (t Taiwan) CreateExchangeRate() IExchangeRate {
	return TWExchangeRate{}
}

func (t Taiwan) CreateStyle() IStyle {
	return TaiwanStyle{}
}
