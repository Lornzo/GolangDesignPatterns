package example1

import "fmt"

type IProduct interface {
	GetName() string
	SetName(name string)
	GetPrice() float64
	SetPrice(price float64)
	GetStock() int
	SetStock(stock int)
	Description()
	Clone() IProduct
}

type AProduct struct {
	Name  string
	Price float64
	Stock int
}

func (p *AProduct) GetName() string {
	return p.Name
}

func (p *AProduct) SetName(name string) {
	p.Name = name
}

func (p *AProduct) GetPrice() float64 {
	return p.Price
}

func (p *AProduct) SetPrice(price float64) {
	p.Price = price
}

func (p *AProduct) GetStock() int {
	return p.Stock
}

func (p *AProduct) SetStock(stock int) {
	p.Stock = stock
}

func (p *AProduct) Description() {
	fmt.Printf("Name: %s, Price: %f, Stock: %d\n", p.Name, p.Price, p.Stock)
}

func (p *AProduct) Clone() IProduct {
	return &AProduct{
		Name:  p.Name,
		Price: p.Price,
		Stock: p.Stock,
	}
}

type Cloth struct {
	IProduct
}

type Pans struct {
	IProduct
}

type Shoes struct {
	IProduct
}
