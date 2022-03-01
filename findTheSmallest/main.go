package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Smallest(261235)) // [126235, 2, 0]
	fmt.Println(Smallest(209917)) // [29917, 0, 1]
	fmt.Println(Smallest(741937037623)) // [74193737623, 7, 1]
	fmt.Println(Smallest(269045)) // [26945, 3, 0]
	fmt.Println(Smallest(44701529626591)) // [4471529626591, 3, 0]
	fmt.Println(Smallest(1000000)) // [1, 0, 6]
}

func Smallest(n int64) []int64 {
	str := strconv.Itoa(int(n))
	var smallest int
	var position int
	strArr := strings.Split(str, "")
	for i, digit := range strArr {
		num, _ := strconv.Atoi(digit)
		if i == 0 {
			smallest = num
			position = i
		} else {
			if smallest >= num {
				smallest = num
				position = i
			}
		}
	}
	// fmt.Println(smallest, position)
	result := []string{}
	result = append(result, strArr[position])
	result = append(result, strArr[:position]...)
	result = append(result, strArr[position+1:]...)
	// fmt.Println("results: ", result)
	format, _ := strconv.ParseInt(strings.Join(result, ""), 10, 64)
	// fmt.Println(format)
	output := []int64{}
	index := 0
	if smallest == 0 {
		position = position + 1 
		index = index + 1
	}
	output = append(output, format, int64(position), int64(index))
	return output
}
