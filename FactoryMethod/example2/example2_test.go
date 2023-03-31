package example2_test

import (
	"fmt"
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/GolangDesignPatterns/FactoryMethod/example2"
)

func TestPlay(t *testing.T) {
	createCar(example2.FamilyCarCreator{})
	createCar(example2.SUVCarCreator{})
	createCar(example2.SportsCarCreator{})
}

func createCar(creator example2.CarCreator) {

	var (
		car example2.Car = creator.CreateCar()
	)

	fmt.Println(car.Interduce())

}
