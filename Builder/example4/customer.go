package example4

type ICustomer interface {
	SetTaste(taste ITaste)
	Order()
}

type CustomerA struct {
	Taste ITaste
}

func (c *CustomerA) SetTaste(taste ITaste) {
	c.Taste = taste
}

func (c *CustomerA) Order() {
	c.Taste.SetBread("wheat")
	c.Taste.SetSize(6)
	c.Taste.Toast()
	c.Taste.AddCheese()
	c.Taste.AddOnion()
	c.Taste.AddOliveOil()
}

type CustomerB struct {
	Taste ITaste
}

func (c *CustomerB) SetTaste(taste ITaste) {
	c.Taste = taste
}

func (c *CustomerB) Order() {
	c.Taste.SetBread("white")
	c.Taste.SetSize(12)
	c.Taste.AddTomato()
	c.Taste.AddLettuce()
	c.Taste.AddThousandIslandDressing()
}
