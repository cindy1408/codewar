package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MultiTable(5))
}

func MultiTable(number int) string {
	table := []string{}
	for i := 1; i <= 10; i++ {
		table = append(table, fmt.Sprintf("%d * %d = %d", i, number, i*number))
	}
	return strings.Join(table, "\n")
}
