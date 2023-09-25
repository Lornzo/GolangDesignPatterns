package example2

type ShapeFactory struct{}

func (s ShapeFactory) CreateCircle(color string, area float64) Shape {
	return NewCircle(color, area)
}

func (s ShapeFactory) CreateTriangle(color string, area float64) Shape {
	return NewTriangle(color, area)
}

func (s ShapeFactory) CreateSquare(color string, area float64) Shape {
	return NewSquare(color, area)
}
