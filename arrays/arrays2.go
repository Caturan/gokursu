package arrays

import "fmt"

func Demo2() {
	var cities [5]string
	cities[0] = "Ankara"
	cities[1] = "Manisa"
	cities[2] = "Izmir"
	cities[3] = "Ağri"
	cities[4] = "Bursa"

	fmt.Println(cities)

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

}
