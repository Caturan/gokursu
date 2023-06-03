package slices

import "fmt"

func Demo1() {
	names := make([]string, 3)

	names[0] = "Ahmet"
	names[1] = "Zeynep"
	names[2] = "Eda"

	fmt.Println(names)
	fmt.Println(len(names))

	names = append(names, "Büşra")

	fmt.Println(names)
	fmt.Println(len(names))

}
