package arrays

import "fmt"

func Demo4() {
	var numbers1 [2][3]int

	numbers1[0][0] = 1
	numbers1[0][1] = 3
	numbers1[0][2] = 5
	numbers1[1][0] = 2
	numbers1[1][1] = 4
	numbers1[1][2] = 6

	fmt.Println(numbers1[1][1])

	for i := 0; i < len(numbers1); i++ {
		for j := 0; j < len(numbers1[i]); j++ {
			fmt.Print(numbers1[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}

}
