package Range

import "fmt"

// Sum of 1-10 odd numbers
// With for-range
func Demo2() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sum := 0

	for _, num := range numbers {
		if num%2 != 0 {
			sum += num
		}
	}
	fmt.Println("Sum : ", sum)
}
