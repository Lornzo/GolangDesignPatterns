package example2

type Shape interface {
	Clone() Shape
	GetColor() string
	SetColor(color string) Shape
	GetArea() float64
	SetArea(area float64) Shape
	PrintDescription()
}
