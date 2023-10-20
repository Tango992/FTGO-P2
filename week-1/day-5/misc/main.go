package main

import "fmt"

type Circle struct {
	Radius float64
	Luas float64
}

type Shape interface {
	Area() 
}

type Test struct{}

func (c *Circle) Area() {
	c.Luas = 3.14 * c.Radius * c.Radius
}

func CalculateArea(shape Shape) {
	shape.Area()
}

func main() {
	circle := Circle{
		Radius: 10.0,
	}

	CalculateArea(&circle)
	fmt.Println(circle)
}
