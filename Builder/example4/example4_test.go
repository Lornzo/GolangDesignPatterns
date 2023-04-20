package example4_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/Builder/example4"
)

func TestExample4(t *testing.T) {

	var (
		customer example4.ICustomer
		chicken  example4.ChickenTaste
		beef     example4.BeefTaste
	)

	customer = &example4.CustomerA{}
	customer.SetTaste(&chicken)
	customer.Order()
	t.Log("Customer A ordered a", chicken.GetSandwich().Description())

	customer = &example4.CustomerB{}
	customer.SetTaste(&beef)
	customer.Order()
	t.Log("Costomer B ordered a", beef.GetSandwich().Description())

}
