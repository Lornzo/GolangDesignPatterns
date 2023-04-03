package example3

import "fmt"

type OrderInfo struct {
	Menu         IMenu
	Style        IStyle
	ExchangeRate IExchangeRate
}

func (o OrderInfo) Print() {
	fmt.Println("Menu: ", o.Menu.GetMenuItems(), "Style: ", o.Style.StyleDescription(), "ExchangeRate: ", o.ExchangeRate.ExchangeRate())
}

type OrderSystem struct {
	Country ICountry
}

func (o OrderSystem) GetOrderInfo() OrderInfo {
	return OrderInfo{
		Menu:         o.Country.CreateMenu(),
		Style:        o.Country.CreateStyle(),
		ExchangeRate: o.Country.CreateExchangeRate(),
	}
}
