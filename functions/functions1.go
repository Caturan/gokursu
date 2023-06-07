package functions

import "fmt"

func Add(num1 int, num2 int) int {
	var addition = num1 + num2
	return addition
}

func SayHello(userName string) {
	fmt.Println("Hello", userName)
}

func Division(num1 int, num2 int) int {
	var division = num1 / num2
	return division
}
