package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	adress             string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	calculater() float64
}

func (m Mortgage) calculater() float64 {
	return m.creditPaymentTotal * m.rate / 12
}

func (c Car) calculater() float64 {
	return c.creditPaymentTotal * c.rate / 6
}

func CalculateMonthlyPayment(credits []CreditCalculator) {
	paymentTotal := 0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + int(credits[i].calculater())
	}

	fmt.Println("Monthly credit payments:", paymentTotal)
}

func Demo3() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 10000, adress: "Izmir"}
	credit2 := Car{rate: 8.5, creditPaymentTotal: 3000, carInfo: "Mercedes"}

	credits := []CreditCalculator{credit1, credit2}
	CalculateMonthlyPayment(credits)
}
