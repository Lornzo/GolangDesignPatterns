package example1

type ProductFactory struct{}

func (p ProductFactory) CreateCloth() *Cloth {
	return &Cloth{&AProduct{}}
}

func (p ProductFactory) CreateShoes() *Shoes {
	return &Shoes{&AProduct{}}
}

func (p ProductFactory) CreatePans() *Pans {
	return &Pans{&AProduct{}}
}
