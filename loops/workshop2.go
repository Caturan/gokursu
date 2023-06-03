package loops

import "fmt"

// 1. Kullanıcıdan bir sayı girmesini isteyiniz.
// Kullanıcının girdiği sayının asal olup olmadığına bakacağız.

func Demo4() {

	var sayi int = 0
	fmt.Println("Bir sayi giriniz: ")

	fmt.Scanln(&sayi)

	var değer int = 0

	for i := 1; i <= sayi; i++ {
		if sayi%i == 0 {
			değer++
		}
	}
	//fmt.Println(değer)

	if değer <= 2 {
		fmt.Println("Asal bir sayidir:)")
	} else {
		fmt.Println("Asal bir sayi değildir:(")
	}

}
