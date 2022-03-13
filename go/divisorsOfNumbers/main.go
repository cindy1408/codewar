package main

import "fmt"

func main() {
	fmt.Println(Divisors(1))
	fmt.Println(Divisors(10))
	fmt.Println(Divisors(11))
	fmt.Println(Divisors(54))
	fmt.Println(Divisors(64))
}

func Divisors(n int) (count int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return
}
