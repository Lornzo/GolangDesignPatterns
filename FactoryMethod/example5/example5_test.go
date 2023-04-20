package example5_test

import (
	"fmt"
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/FactoryMethod/example5"
)

func TestExample5(t *testing.T) {

	painKillerFactory := &example5.PainKillerFactory{}
	antipyreticsFactory := &example5.AntipyreticsFactory{}
	antiInflammatoryFactory := &example5.AntiInflammatoryFactory{}

	painKiller := painKillerFactory.MakeMedicine()
	antipyretics := antipyreticsFactory.MakeMedicine()
	antiInflammatory := antiInflammatoryFactory.MakeMedicine()

	useMedicine(painKiller)
	useMedicine(antipyretics)
	useMedicine(antiInflammatory)
}

func useMedicine(medicine example5.Medicine) {
	fmt.Println(medicine.GetIngredients())
}
