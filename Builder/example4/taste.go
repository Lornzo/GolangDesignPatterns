package example4

type ITaste interface {
	SetBread(bread string)
	SetSize(size int)
	Toast()
	AddCheese()
	AddOnion()
	AddTomato()
	AddLettuce()
	AddOliveOil()
	AddThousandIslandDressing()
}

type ChickenTaste struct {
	Sandwich ChickenSandwich
}

func (c *ChickenTaste) SetBread(bread string) {
	c.Sandwich.Bread = bread
}

func (c *ChickenTaste) SetSize(size int) {
	c.Sandwich.Size = size
}

func (c *ChickenTaste) Toast() {
	c.Sandwich.Toasted = true
}

func (c *ChickenTaste) AddCheese() {
	c.Sandwich.Recipt = append(c.Sandwich.Recipt, "cheese")
}

func (c *ChickenTaste) AddOnion() {
	c.Sandwich.Recipt = append(c.Sandwich.Recipt, "onion")
}

func (c *ChickenTaste) AddTomato() {
	c.Sandwich.Recipt = append(c.Sandwich.Recipt, "tomato")
}

func (c *ChickenTaste) AddLettuce() {
	c.Sandwich.Recipt = append(c.Sandwich.Recipt, "lettuce")
}

func (c *ChickenTaste) AddOliveOil() {
	c.Sandwich.Recipt = append(c.Sandwich.Recipt, "olive oil")
}

func (c *ChickenTaste) AddThousandIslandDressing() {
	c.Sandwich.Recipt = append(c.Sandwich.Recipt, "thousand island dressing")
}

func (c *ChickenTaste) GetSandwich() ChickenSandwich {
	return c.Sandwich
}

type BeefTaste struct {
	Sandwich BeefSandwich
}

func (b *BeefTaste) SetBread(bread string) {
	b.Sandwich.Bread = bread
}

func (b *BeefTaste) SetSize(size int) {
	b.Sandwich.Size = size
}

func (b *BeefTaste) Toast() {
	b.Sandwich.Toasted = true
}

func (b *BeefTaste) AddCheese() {
	b.Sandwich.Recipt = append(b.Sandwich.Recipt, "cheese")
}

func (b *BeefTaste) AddOnion() {
	b.Sandwich.Recipt = append(b.Sandwich.Recipt, "onion")
}

func (b *BeefTaste) AddTomato() {
	b.Sandwich.Recipt = append(b.Sandwich.Recipt, "tomato")
}

func (b *BeefTaste) AddLettuce() {
	b.Sandwich.Recipt = append(b.Sandwich.Recipt, "lettuce")
}

func (b *BeefTaste) AddOliveOil() {
	b.Sandwich.Recipt = append(b.Sandwich.Recipt, "olive oil")
}

func (b *BeefTaste) AddThousandIslandDressing() {
	b.Sandwich.Recipt = append(b.Sandwich.Recipt, "thousand island dressing")
}

func (b *BeefTaste) GetSandwich() BeefSandwich {
	return b.Sandwich
}
