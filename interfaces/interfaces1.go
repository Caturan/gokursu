package interfaces

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func school(s Shape) {
	fmt.Println("Shape of area: ", s.area())
}

func Demo1() {
	r1 := Rectangle{width: 10, height: 3}
	school(r1)

	c1 := Circle{radius: 8}
	school(c1)
}
