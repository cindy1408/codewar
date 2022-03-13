package main

import "fmt"

func main() {
	fmt.Println(GetSize(4, 2, 6))
	fmt.Println(GetSize(1, 1, 1))
	fmt.Println(GetSize(1, 2, 1))
	fmt.Println(GetSize(1, 2, 2))
	fmt.Println(GetSize(10, 10, 10))
}

// [area, volume]
func GetSize(w, h, d int) (sizes [2]int) {
	sizes[0] = 2 * ((w * h) + (w * d) + (h * d))
	sizes[1] = w * h * d
	return
}
