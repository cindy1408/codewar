package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(NbDig(10, 1))
	fmt.Println(NbDig(25, 1))
	fmt.Println(NbDig(550, 5))
	fmt.Println(NbDig(5750, 0))
}

func NbDig(n int, d int) int {
	listOfSquares := []int{}
	for i := 0; i <= n; i++ {
		listOfSquares = append(listOfSquares, i*i)
	}
	totalNums := 0
	for _, num := range listOfSquares {
		for _, digit := range strings.Split(strconv.Itoa(int(num)), "") {
			digitNum, _ := strconv.Atoi(digit)
			if digitNum == d {
				totalNums++
			}
		}
	}
	return totalNums
}
