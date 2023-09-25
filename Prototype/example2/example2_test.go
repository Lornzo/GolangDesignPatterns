package example2

import "testing"

func TestExample2(t *testing.T) {
	var factory ShapeFactory = ShapeFactory{}

	var circle Shape = factory.CreateCircle("Red", 10)
	circle.PrintDescription()

	var anotherCircle Shape = circle.Clone()
	anotherCircle.SetColor("Blue")
	anotherCircle.PrintDescription()
}
