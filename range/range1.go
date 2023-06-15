package Range

import "fmt"

func Demo1() {
	cities := []string{"Berlin", "NewYork", "London"}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	for _, city := range cities {
		fmt.Println(city)
	}

}
