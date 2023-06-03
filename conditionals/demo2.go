package conditionals

import "fmt"

func Demo2() {
	hesap := 1500
	cekilmekistenen := 1200

	hesap1 := hesap - cekilmekistenen

	if hesap > cekilmekistenen {
		fmt.Println("Paraniz hazirlaniyor")
		fmt.Println("Hesapta kalan paraniz : " + fmt.Sprintf("%d", hesap1))
	} else if hesap == cekilmekistenen {
		fmt.Println("Paraniz hazirlaniyor")
		fmt.Println("Hesapta kalan paraniz : " + fmt.Sprintf("%d", hesap1))
	} else {
		fmt.Println("Hesaptaki para yetersiz")
	}

}
