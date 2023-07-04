package deferstatement

import "fmt"

type Product struct {
	ProductName string
	UnitPrice   int
}

func (p Product) save() {
	fmt.Println("Product's name saved: ", p.ProductName)
	defer Log()
}

func (p Product) save2() {
	fmt.Println("Product's price saved: ", p.UnitPrice)
}

func Log() {
	fmt.Println("Logged")
}

func Demo3() {
	k := Product{ProductName: "Apple", UnitPrice: 999}
	k.save()
	defer k.save2()
	fmt.Println("Finished")

}
