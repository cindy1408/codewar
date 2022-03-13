package main

import "fmt"

func main() {
	fmt.Println(Race(720, 850, 70))
	fmt.Println(Race(820, 81, 550))
	fmt.Println(Race(80, 91, 37))
}

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	time := (g * 3600)/(v2-v1)
	return [3]int{time/3600, time%3600/60, time%3600%60}
}
