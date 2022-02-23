package main

import (
	"fmt"
	"math"
	"regexp"
)

func main() {
	fmt.Println(ListSquared(1, 246))
	fmt.Println(ListSquared(1, 250))
	fmt.Println(ListSquared(250, 500))
	fmt.Println(ListSquared(300, 600))
	fmt.Println(ListSquared(4099, 4999))
}

func ListSquared(m, n int) [][]int {
	result := [][]int{}
	for i := m; i <= n; i++ {
		fArr := divisors(i)
		total := 0
		for _, num := range fArr {
			total = total + int(math.Pow(float64(num), float64(2)))
		}
		reg := regexp.MustCompile(`\.(0+$)`)
		integer := reg.MatchString(fmt.Sprintf("%f", math.Sqrt(float64(total))))
		if integer {
			result = append(result, []int{i, total})
		}
	}
	return result
}

func divisors(num int) (result []int) {
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			result = append(result, i)
		}
	}
	return
}
