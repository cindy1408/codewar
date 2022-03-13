package main

import (
	"fmt"
)

func main() {
	fmt.Print(Amort(6, 100000, 360, 1)) //"num_payment 1 c 600 princ 100 int 500 balance 99900"

	fmt.Print(Amort(6, 100000, 360, 12)) //"num_payment 12 c 600 princ 105 int 494 balance 98772"
}

func Amort(rate float64, balance int, term int, num_payments int) string {

	monthlyRate := (rate/100) * 12
	// n := rate * float64(balance)
	// d := math.Pow(1 - (1+rate),-float64(term))
	// c is the payment = interests + princ
	interestPerMonth := monthlyRate * float64(balance)
	c := float64(balance)/12 + interestPerMonth
	princ := float64(balance) / float64(num_payments)
	interest := (monthlyRate * float64(balance)) / float64(num_payments) 
	remainBalance := float64(balance) - princ
	return fmt.Sprintf("num_payment %d c %.0f princ %.0f int %.0f balance %.0f\n", num_payments, c, princ, interest, remainBalance)
}
