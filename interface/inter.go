package main

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

func main() {
	square := Square{5}
	// circle := Circle{7}

	PrintShapeArea(square)
	// PrintShapeArea(circle)

	// printInterface(square)
	// printInterface(circle)
	printInterface("sada")
	printInterface(77)
	printInterface("ggtdvdv")
	// printInterface(true)
}

func PrintShapeArea(shape Shape) {
	fmt.Println(shape.Area())
	fmt.Println(shape.Perimetr())
}

func printInterface (i interface{}) {
// 	switch value := i.(type) {
// 	case int: 
// 		fmt.Println("int", value)
// 	case bool: 
// 		fmt.Println("bool", value)
// 	default: 
// 		fmt.Println("unknown type", value)
// 	}
// 	fmt.Printf("%#v\n", i)

	str, ok := i.(string)
	if !ok {
		fmt.Println("interface is not string")
		return
	}
	fmt.Println(len(str))
}