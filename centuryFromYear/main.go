package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Century(1990))
	fmt.Println(Century(1705))
	fmt.Println(Century(1900))
	fmt.Println(Century(1601))
	fmt.Println(Century(2000))
	fmt.Println(Century(89))
}

func Century(year int) int {
	return int(math.Ceil(float64(year)/100))
}
