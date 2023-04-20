package example2_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/Builder/example2"
)

func TestExample2(t *testing.T) {

	var (
		teaBuilder   example2.WooLongTeaBuilder
		juiceBuilder example2.AppleJuiceBuilder
		maker        example2.IBeverageMaker
		beverage     example2.IBeverage
	)

	maker = &example2.HelfSugerMaker{}
	maker.SetBuilder(&juiceBuilder)
	maker.Make()
	beverage = juiceBuilder.GetBaverage()
	t.Log(beverage.Description())

	maker = &example2.HighTemperatureMaker{}
	maker.SetBuilder(&teaBuilder)
	maker.Make()
	beverage = teaBuilder.GetBaverage()
	t.Log(beverage.Description())
}
