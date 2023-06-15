package maps

import "fmt"

func Demo1() {
	// key-value
	dictionary := make(map[string]string)

	dictionary["Apple"] = "Elma"
	dictionary["Book"] = "Kitap"
	dictionary["Chair"] = "Sandalye"
	dictionary["Chief"] = "Şef"

	fmt.Println(dictionary["Chair"])
	fmt.Println("Dictionary :", dictionary)
	fmt.Println("Number of elements:", len(dictionary))
	delete(dictionary, "Apple")
	fmt.Println("Dictionary :", dictionary)
	fmt.Println("Number of elements:", len(dictionary))

	value, isIt := dictionary["Apple"]
	fmt.Println(value)
	fmt.Println("In the list: ", isIt)

	fmt.Println()

	dictionary2 := map[string]string{"Keyboard": "Klavye", "Bottle": "Şişe"}
	fmt.Println(dictionary2)

}
