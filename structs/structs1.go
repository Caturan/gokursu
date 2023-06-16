package structs

import "fmt"

func Demo1() {

	fmt.Println(product{"iPhone", 599, "Apple"})
	x := product{"Note", 350, "Samsung"}
	fmt.Println(x)
	fmt.Println(product{name: "xPreia"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
