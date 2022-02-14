package main

import (
	"strings"
)

func main() {
	s1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
	s2 := []string{"bbbaaayddqbbrrrv"}
	MxDifLg(s1, s2)
}

func MxDifLg(a1 []string, a2 []string) int {
	if len(a1) == 0 || len(a2) == 0 {
		return -1
	}
	m := make(map[string][]int)
	for _, str := range a1 {
		m["a1"] = append(m["a1"], len(strings.Split(str, "")))
	}
	for _, str := range a2 {
		m["a2"] = append(m["a2"], len(strings.Split(str, "")))
	}
	var min int
	var max int
	maxArr := []int{}
	minArr := []int{}
	for _, numArr := range m {
		for i, num := range numArr {
			if i == 0 {
				min = num
				max = num
			}
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		maxArr = append(maxArr, max)
		minArr = append(minArr, min)
	}
	if maxArr[0]-minArr[1] > maxArr[1]-minArr[0] {
		return maxArr[0] - minArr[1]
	} else {
		return maxArr[1] - minArr[0]
	}
}
