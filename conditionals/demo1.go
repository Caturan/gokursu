package conditionals

import "fmt"

func Demo1() {
	hesap := 1000
	cekilmekistenen := 900

	hesap1 := hesap - cekilmekistenen
	hesap2 := hesap - cekilmekistenen

	if cekilmekistenen > hesap {
		fmt.Println("Para Yetersiz")
		fmt.Println("Hesaptaki para : " + fmt.Sprintf("%d", hesap1))
	}

	if cekilmekistenen <= hesap {
		fmt.Println("Paraniz hazirlaniyor")
		fmt.Printf("Hesaptaki para : %d", hesap2)
	}

}
