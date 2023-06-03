package slices

import "fmt"

func Demo2() {
	cities := []string{"Ankara, Istanbul", "Diyarbakir"}
	fmt.Println(cities)

	citiesCopy := make([]string, len(cities))
	fmt.Println(citiesCopy)

	copy(citiesCopy, cities)
	fmt.Println(citiesCopy)

	fmt.Println()

	cities = append(cities, "Izmir")
	fmt.Println(cities)
	fmt.Println(citiesCopy)

	fmt.Println()

	fmt.Println(cities[1:3])
	fmt.Println(cities[2:])
	fmt.Println(cities[:2])
}
