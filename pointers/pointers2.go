package pointers

import "fmt"

func Demo2(numbers []int) {
	numbers[0] = 100
	fmt.Println("Result in demo :", numbers[0])
}

// Arrays doesn't value type
