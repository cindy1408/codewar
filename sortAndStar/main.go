package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(TwoSort([]string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}))
	fmt.Println(TwoSort([]string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}))
	fmt.Println(TwoSort([]string{"lets", "talk", "about", "javascript", "the", "best", "language"}))
	fmt.Println(TwoSort([]string{"i", "want", "to", "travel", "the", "world", "writing", "code", "one", "day"}))
	fmt.Println(TwoSort([]string{"Lets", "all", "go", "on", "holiday", "somewhere", "very", "cold"}))
}

func TwoSort(arr []string) string {
	sort.Strings(arr)
	chosen := strings.Split(arr[0], "") 
	return strings.Join(chosen, "***")
}
