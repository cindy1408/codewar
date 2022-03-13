package main

import "fmt"

func main() {
	fmt.Println(Summation(1))
	fmt.Println(Summation(8))
	fmt.Println(Summation(22))
	fmt.Println(Summation(100))
	fmt.Println(Summation(213))
}

func Summation(n int) (summation int) {
	for i := 0; i <= n; i++ {
		summation = summation + i
	}
	return
}
