package loops

import "fmt"

// Arkadaş sayılar problemi
// Kendisi hariç bölenlerinin toplamı birbirine eşit olan sayılar.
// Örnek: 220 ve 284
// 220 : 1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110 = 284 284 : 1 + 2 + 4 + 71 + 142 = 220

func Demo5() {

	var sayi1 int = 0
	var sayi2 int = 0

	fmt.Println("Lütfen 2 tane sayi giriniz: ")
	fmt.Scanln(&sayi1)
	fmt.Scanln(&sayi2)

	var toplam1 int = 0
	var toplam2 int = 0

	for i := 1; i < sayi1; i++ {
		if sayi1%i == 0 {
			toplam1 += i
		}
	}
	fmt.Println(toplam1)

	for i := 1; i < sayi2; i++ {
		if sayi2%i == 0 {
			toplam2 += i
		}
	}
	fmt.Println(toplam2)

	if toplam1 == sayi2 && toplam2 == sayi1 {
		fmt.Println("Bu sayilar kardeş sayidir:)")
	} else {
		fmt.Println("Değildir:(")
	}

}
