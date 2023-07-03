package channels

import (
	"fmt"
	"time"
)

func EvenNumbers(evenNumbersCn chan int) {
	sum := 0
	for i := 0; i <= 10; i += 2 {
		sum += i
		fmt.Println("Even numbers is working")
		time.Sleep(time.Second)
	}
	evenNumbersCn <- sum
}

func OddNumbers(oddNumbersCn chan int) {
	sum := 0
	for i := 1; i <= 10; i += 2 {
		sum += i
		fmt.Println("Odd numbers is working")
		time.Sleep(time.Second)
	}
	oddNumbersCn <- sum
}
