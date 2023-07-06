package errorhandling

import (
	"fmt"
	"os"
)

// Type Assertion
func Demo1() {
	f, err := os.Open("Demo2.txt")
	//nil
	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("File not found :", pErr.Path)
			return
		} else {
			fmt.Println("File not found")
			return
		}
	} else {
		fmt.Println(f.Name())
	}

}
