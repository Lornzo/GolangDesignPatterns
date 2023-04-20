package example1_test

import (
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/Builder/example1"
)

func TestExample1(t *testing.T) {

	var (
		house    example1.IHouse
		director example1.IDirector
	)

	strawBuilder := example1.StrawHouseBuilder{}
	director = &example1.SimpleHouseDirector{}
	director.SetBuilder(&strawBuilder)
	director.LetBuilderBuildingHouse()
	house = strawBuilder.GetHouse()
	t.Log(house.Interduce())

	woodBuilder := example1.WoodHouseBuilder{}
	director = &example1.LuxuryHouseDirector{}
	director.SetBuilder(&woodBuilder)
	director.LetBuilderBuildingHouse()
	house = woodBuilder.GetHouse()
	t.Log(house.Interduce())

}
