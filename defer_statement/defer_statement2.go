package deferstatement

import "fmt"

func Demo2(num int32) string {
	defer DeferFunc()

	if num%2 == 0 {
		return "Even number"
	}

	if num%2 != 0 {
		fmt.Println("Odd number executed")
		return "Odd number"
	}
	return "Default"
}

func Test() {
	fmt.Println(Demo2(15))
}

func DeferFunc() {
	fmt.Println("Finished")
}
