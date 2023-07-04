package deferstatement

import (
	"fmt"
)

func A() {
	fmt.Println("Func A executed")
}

func B() {
	defer A() // LIFO(Last In First Off)
	defer C()
	defer D()
	fmt.Println("Func B executed")
	A()
}

func C() {
	fmt.Println("Func C executed")
}

func D() {
	fmt.Println("Func D executed")
}
