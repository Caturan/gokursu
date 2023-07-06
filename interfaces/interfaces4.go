package interfaces

import "fmt"

func Verify(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo4() {

	Verify("TURAN")

	Verify(7)
}
