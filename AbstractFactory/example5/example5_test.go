package example5_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/AbstractFactory/example5"
)

func TestExample5(t *testing.T) {
	var store example5.BagStore
	store.Factory = example5.LuxuryBagFactory{}
	store.ShowAllBags()

	store.Factory = example5.CasualBagFactory{}
	store.ShowAllBags()
}
