package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsTriangle(5, 1, 2))
	fmt.Println(IsTriangle(1, 2, 5))
	fmt.Println(IsTriangle(2, 5, 1))
	fmt.Println(IsTriangle(4, 2, 3))
	fmt.Println(IsTriangle(5, 1, 5))
	fmt.Println(IsTriangle(2, 2, 2))
	fmt.Println(IsTriangle(-1, 2, 3))

}

func IsTriangle(a, b, c int) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	if a+b > c && b+c > a && a+c > b {
		return true
	}
	return false
}
