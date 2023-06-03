package loops

import "fmt"

func Demo3() {

	var AklimdakiSayi int = 0
	var TahminEdilenSayi int = 0

	fmt.Println("Aklinizdaki sayiyi giriniz: ")
	fmt.Scanln(&AklimdakiSayi)

	/*
		fmt.Println("INPUT: ")
		fmt.Scanln(&TahminEdilenSayi)
		fmt.Println(TahminEdilenSayi)
	*/
	//  Sayı tahmin oyunu
	counter := 1

	for j := 0; j < 100; j++ {
		fmt.Println("INPUT: ")
		fmt.Scanln(&TahminEdilenSayi)
		if AklimdakiSayi > TahminEdilenSayi {
			fmt.Println("UP!")
		} else if AklimdakiSayi < TahminEdilenSayi {
			fmt.Println("DOWN!")
		} else if AklimdakiSayi == TahminEdilenSayi {
			fmt.Println("Correct! ", AklimdakiSayi)
			break
		}
		counter++

	}
	if counter <= 3 {
		fmt.Println("Bravo", counter, ". seferinizde bildiniz. Perfect match-up :)")
	} else if 7 >= counter && counter > 3 {
		fmt.Println("İyi", counter, ". seferinizde bildiniz. Eh işte :|")
	} else if counter > 7 {
		fmt.Println("Kötü", counter, ". seferinizde bildiniz. Uyum negatif:(")
	}

}
