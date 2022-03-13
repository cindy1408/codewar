package main

import "fmt"

func main() {
	fmt.Println(CheckForFactor(63, 7))
	fmt.Println(CheckForFactor(10, 2))
	fmt.Println(CheckForFactor(2450, 5))
	fmt.Println(CheckForFactor(24612, 3))
	fmt.Println(CheckForFactor(9, 2))
	fmt.Println(CheckForFactor(653, 7))
	fmt.Println(CheckForFactor(2453, 5))
	fmt.Println(CheckForFactor(24617, 3))
}

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0;
}
