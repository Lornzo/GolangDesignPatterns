package example3_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/AbstractFactory/example3"
)

func TestExample3(t *testing.T) {

	var (
		client example3.OrderSystem
		info   example3.OrderInfo
	)

	client.Country = example3.Taiwan{}
	info = client.GetOrderInfo()
	info.Print()

	client.Country = example3.USA{}
	info = client.GetOrderInfo()
	info.Print()

}
