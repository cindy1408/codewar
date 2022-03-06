package main

import "fmt"

func main() {
	fmt.Println(MultiplicationTable(1)) //{1}
	fmt.Println(MultiplicationTable(2)) //{1, 2}, {2, 4}
	fmt.Println(MultiplicationTable(3)) //{1, 2, 3}, {2, 4, 6}, {3, 6, 9}
}

func MultiplicationTable(size int) [][]int {
	table := [][]int{}
	for i := 1; i <= size; i++ {
		eachMultiple := []int{}
		for j := 1; j <= size; j++ {
			eachMultiple = append(eachMultiple, i*j)
		}
		table = append(table, eachMultiple)
	}
	return table
}
