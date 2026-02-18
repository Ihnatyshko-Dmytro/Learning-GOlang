package shape

import (
	"fmt"
	"math"
)

type ShapeWithArea interface {
	Area() float32
}

type ShapeWithPerimeter interface {
	Perimetr() float32
}

type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

type Square struct {
	sideLenght float32
}

func NewSquare(lenght float32) Square {
	return Square{
		sideLenght: lenght,
	}
}

func (s Square) Area() float32 {
	return s.sideLenght * s.sideLenght
}

func (s Square) Perimetr() float32 {
	return s.sideLenght * 4
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func PrintShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimetr())
}
