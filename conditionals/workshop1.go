package conditionals

import "fmt"

func Demo3() {
	// Üç adet int değişkeni
	// Ekrana en büyük olanı bastır.

	var1, var2, var3 := 13, 04, 19

	if var1 > var2 {
		if var1 > var3 {
			fmt.Println("En büyük deger var1 : " + fmt.Sprintf("%d", var1))
		}
	}
	if var2 > var1 {
		if var2 > var3 {
			fmt.Println("En büyük deger var2 : " + fmt.Sprintf("%d", var2))
		}
	}
	if var3 > var1 {
		if var3 > var2 {
			fmt.Println("En büyük deger var3 : " + fmt.Sprintf("%d", var3))
		}
	}
}
