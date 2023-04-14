package example2

type IBeverageBuilder interface {
	SetConcertraction(concertraction int)
	SetSugar(sugar int)
	SetMilk(milk int)
	SetTemperature(temp float64)
	SetCarbonic(carbonic int)
}

type CoffeeBuilder struct {
	Baverage Coffee
}

func (c *CoffeeBuilder) SetConcertraction(concertraction int) {
	c.Baverage.Concentration = concertraction
}

func (c *CoffeeBuilder) SetSugar(sugar int) {
	c.Baverage.Sugar = sugar
}

func (c *CoffeeBuilder) SetMilk(milk int) {
	c.Baverage.Milk = milk
}

func (c *CoffeeBuilder) SetTemperature(temp float64) {
}

func (c *CoffeeBuilder) SetCarbonic(carbonic int) {
}

func (c *CoffeeBuilder) GetBaverage() Coffee {
	return c.Baverage
}

type WooLongTeaBuilder struct {
	Beverage Tea
}

func (w *WooLongTeaBuilder) SetConcertraction(concertraction int) {
}

func (w *WooLongTeaBuilder) SetSugar(sugar int) {
	w.Beverage.Sugar = sugar
}

func (w *WooLongTeaBuilder) SetMilk(milk int) {
}

func (w *WooLongTeaBuilder) SetTemperature(temp float64) {
	w.Beverage.Temperature = temp
}

func (w *WooLongTeaBuilder) SetCarbonic(carbonic int) {
}

func (w *WooLongTeaBuilder) GetBaverage() Tea {
	var beverage Tea = w.Beverage
	beverage.TeaType = "WooLong"
	return beverage
}

type AppleJuiceBuilder struct {
	Beverage Juice
}

func (a *AppleJuiceBuilder) SetConcertraction(concertraction int) {
}

func (a *AppleJuiceBuilder) SetSugar(sugar int) {
	a.Beverage.Sugar = sugar / 2
}

func (a *AppleJuiceBuilder) SetMilk(milk int) {
}

func (a *AppleJuiceBuilder) SetTemperature(temp float64) {
}

func (a *AppleJuiceBuilder) SetCarbonic(carbonic int) {
}

func (a *AppleJuiceBuilder) GetBaverage() Juice {
	var beverage Juice = a.Beverage
	beverage.Fruit = "Apple"
	return beverage
}

type ColaSodaBuilder struct {
	Beverage Soda
}

func (c *ColaSodaBuilder) SetConcertraction(concertraction int) {
}

func (c *ColaSodaBuilder) SetSugar(sugar int) {
}

func (c *ColaSodaBuilder) SetMilk(milk int) {
}

func (c *ColaSodaBuilder) SetTemperature(temp float64) {
}

func (c *ColaSodaBuilder) SetCarbonic(carbonic int) {
	c.Beverage.Carbonic = carbonic
}

func (c *ColaSodaBuilder) GetBaverage() Soda {
	var beverage Soda = c.Beverage
	beverage.Style = "Cola"
	return beverage
}
