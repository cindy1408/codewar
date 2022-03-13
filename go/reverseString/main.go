package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Solution("world"))
}

func Solution(word string) string {
	strArr := strings.Split(word, "")
	reverseArr := []string{}
	for i := len(strArr) - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, strArr[i])
	}
	return strings.Join(reverseArr, "")
}

// better solution

// func Solution(word string) string {
// 	var result = ""
// 	for _, c := range word {
// 		result = string(c) + result
// 	}
// 	return result
// }
