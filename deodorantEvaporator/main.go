package main

import "fmt"

func main() {
	fmt.Println(Evaporator(10, 10, 10))
	fmt.Println(Evaporator(10, 10, 5))
	fmt.Println(Evaporator(100, 5, 5))
}

func Evaporator(content float64, evapPerDay int, threshold int) int {
	remainder := content
	thresholdml := float64(threshold) * 0.01 * content
	days := 0
	for {
		remainder = remainder - (remainder*float64(evapPerDay) * 0.01)
		days++
		if remainder == thresholdml || remainder < thresholdml{
			break
		}
	}
	return days
}