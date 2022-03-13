package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortNumbers([]int{4, 9, 1, 5}))
}

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}
