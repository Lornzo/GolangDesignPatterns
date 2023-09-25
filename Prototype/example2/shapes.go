package example2

import "fmt"

type shape struct {
	color string
	area  float64
}

func (s *shape) Clone() Shape {
	return &shape{
		color: s.color,
		area:  s.area,
	}
}

func (s *shape) GetColor() string {
	return s.color
}

func (s *shape) SetColor(color string) Shape {
	s.color = color
	return s
}

func (s *shape) GetArea() float64 {
	return s.area
}

func (s *shape) SetArea(area float64) Shape {
	s.area = area
	return s
}

func (s *shape) PrintDescription() {
	fmt.Print("Color:", s.color, " | ", "Area:", s.area)
}

func NewCircle(color string, area float64) Shape {
	return &Circle{Shape: &shape{color: color, area: area}}
}

type Circle struct {
	Shape
}

func (c *Circle) Clone() Shape {
	return &Circle{Shape: c.Shape.Clone()}
}

func (c *Circle) PrintDescription() {
	fmt.Print("Circle - ")
	c.Shape.PrintDescription()
	fmt.Print("\n")
}

func NewTriangle(color string, area float64) Shape {
	return &Triangle{Shape: &shape{color: color, area: area}}
}

type Triangle struct {
	Shape
}

func (t *Triangle) Clone() Shape {
	return &Triangle{Shape: t.Shape.Clone()}
}

func (t *Triangle) PrintDescription() {
	fmt.Print("Triangle - ")
	t.Shape.PrintDescription()
	fmt.Print("\n")
}

func NewSquare(color string, area float64) Shape {
	return &Square{Shape: &shape{color: color, area: area}}
}

type Square struct {
	Shape
}

func (s *Square) Clone() Shape {
	return &Square{Shape: s.Shape.Clone()}
}

func (s *Square) PrintDescription() {
	fmt.Print("Square - ")
	s.Shape.PrintDescription()
	fmt.Print("\n")
}
