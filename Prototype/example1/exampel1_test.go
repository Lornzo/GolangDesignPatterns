package example1

import "testing"

func TestExample1(t *testing.T) {

	var (
		product IProduct
		factory interface {
			CreateCloth() *Cloth
			CreateShoes() *Shoes
			CreatePans() *Pans
		} = ProductFactory{}
	)

	product = factory.CreateCloth()
	product.SetName("cloth")
	product.SetPrice(100)
	product.SetStock(10)
	product.Description()

	product = factory.CreateShoes()
	product.SetName("shoes")
	product.SetPrice(200)
	product.SetStock(20)
	product.Description()

	product = product.Clone()
	product.SetName("shoes2")
	product.SetPrice(300)
	product.SetStock(30)
	product.Description()

}
