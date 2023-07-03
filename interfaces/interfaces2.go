package interfaces

import "fmt"

type Shape2 interface {
	perimeter() float64
}

type Squeare struct {
	edge float64
}

type Trieangle struct {
	edge1, edge2, edge3 float64
}

func (s Squeare) perimeter() float64 {
	return s.edge * 4
}

func (t Trieangle) perimeter() float64 {
	return t.edge1 + t.edge2 + t.edge3
}

func school2(s Shape2) {
	fmt.Println("Shape of perimeter:", s.perimeter())
}

func Demo2() {
	s1 := Squeare{19}
	school2(s1)

	t1 := Trieangle{7, 6, 3}
	school2(t1)
}
