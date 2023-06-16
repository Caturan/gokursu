package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (k customer) save() {
	fmt.Println("Added : ", k.firstName, k.lastName)
}

func Demo2() {
	c := customer{"Bob", "Marley", 86}
	c.save()
	fmt.Println("Age is: ", c.age)
}
