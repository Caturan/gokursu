package arrays

import "fmt"

func Demo3() {
	numbers := [5]int{66, 32, 74, 35, 87}

	fmt.Println(numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	hiNumb := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if numbers[i] > hiNumb {
			hiNumb = numbers[i]
		}
	}
	fmt.Println("High number is: ", hiNumb)
}
