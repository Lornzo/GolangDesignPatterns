package example2_test

import (
	"fmt"
	"testing"

	"github.com/Lornzo/GolangDesignPatterns/GolangDesignPatterns/AbstractFactory/example2"
)

type Car struct {
	Type   string
	Engine example2.ICarEngine
	Body   example2.ICarBody
	Wheel  example2.ICarWheel
}

func (c Car) Run() {
	fmt.Println(c.Type, "啟動：", c.Engine.Run(), c.Body.Description(), c.Wheel.Description())
}

type AssemblyFactory struct {
	Factory example2.ICarFactory
}

func (a AssemblyFactory) CreateCar() Car {
	return Car{
		Type:   a.Factory.GetProduceType(),
		Engine: a.Factory.CreateEngine(),
		Body:   a.Factory.CreateBody(),
		Wheel:  a.Factory.CreateWheel(),
	}
}

func TestExample2(t *testing.T) {

	var (
		assembler AssemblyFactory
		car       Car
	)

	assembler.Factory = example2.VanFactory{}
	car = assembler.CreateCar()
	car.Run()

	assembler.Factory = example2.SportCarFactory{}
	car = assembler.CreateCar()
	car.Run()

}
