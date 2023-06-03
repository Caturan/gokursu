package variables

import "fmt"

// camelCase
// PascalCase

func Demo1() {
	var metin string = "Hello World"
	fmt.Println(metin)
	fmt.Println("Lütfen çaliş")
	fmt.Println("Teşekkürler:)")

	var kdv int = 14
	fmt.Println(kdv)
	fmt.Println(100 + (kdv * 100 / 100))
	fmt.Println()

	var kdv2 float32 = 0.14
	fmt.Println(kdv2)
	fmt.Println(100 + (kdv2 * 100 / 100))
	fmt.Println()

	kdv3 := 25
	fmt.Println(kdv3)
	fmt.Printf("veri tipi : %T\n", kdv3)

	fmt.Println()
	var durum bool
	var durum2 bool

	var metin1 string = "Engin"
	var metin2 string = "Ahmet"

	durum = (metin1 == metin2)
	fmt.Println(durum)

	fmt.Println()

	durum2 = (metin1 != metin2)
	fmt.Println(durum2)
}
