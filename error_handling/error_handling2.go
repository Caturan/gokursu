package errorhandling

import (
	"errors"
	"fmt"
)

func Quess(quess int) (string, error) {

	myNumber := 19

	if quess < 1 || quess > 100 {
		return "", errors.New("Input a number between 1 - 100")
	}

	if quess > myNumber {
		return "Input a bigger number", nil
	}

	if quess < myNumber {
		return "Input a smaller number", nil
	}

	return "Correct", nil
}

func Demo2() {
	fmt.Println(Quess(19))
	fmt.Println(Quess(101))
	fmt.Println(Quess(-2))

	message, _ := Quess(23)
	fmt.Println(message)
	//fmt.Println(fail)

}
