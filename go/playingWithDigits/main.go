package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(DigPow(89, 1))
	fmt.Println(DigPow(92, 1))
}

func DigPow(n, p int) int {
	total := 0.00
	for i, digit := range strings.Split(strconv.Itoa(n), "") {
		d, _ := strconv.Atoi(digit)
		total = total + math.Pow(float64(d), float64(p+i))
	}
	if int(total) % n != 0 {
		return -1
	}
	return int(total)/n
}
