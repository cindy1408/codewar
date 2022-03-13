package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(duplicate_count("abcde"))
	fmt.Println(duplicate_count("abcdea"))
	fmt.Println(duplicate_count("abcdeaB11"))
	fmt.Println(duplicate_count("indivisibility"))
}

func duplicate_count(s1 string) int {
	var count int
	tracker := make(map[string]int)
	for _, char := range strings.Split(s1, "") {
		tracker[strings.ToLower(char)] = tracker[strings.ToLower(char)] + 1 
	}
	for _, occurance := range tracker {
		if occurance != 1 {
			count++
		}
	}
	return count
}
