package main

import (
	"strings"
)

func main() {
	vowelCount("hello")
}

func vowelCount(str string) int {
	stringArr := strings.Split(str, "")
	count := 0
	for _, char := range stringArr {
		switch char {
		case "a", "e", "i", "o", "u":
			count++
		}
	}
	return count
}
