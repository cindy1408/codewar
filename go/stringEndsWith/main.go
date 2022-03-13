package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(solution("", ""))
	fmt.Println(solution(" ", ""))
	fmt.Println(solution("abc", "c"))
	fmt.Println(solution("banana", "ana"))
	fmt.Println(solution("a", "z"))
	fmt.Println(solution("", "t"))
	fmt.Println(solution("$a = $b + 1", "+1"))
	fmt.Println(solution("    ", "   "))
	fmt.Println(solution("1oo", "100"))
}

func solution(str, ending string) bool {
	strArr := strings.Split(str, "")
	endingArr := strings.Split(ending, "")
	matchEndingArr := []string{}
	if len(strArr) == 0 && len(endingArr) != 0 {
		return false
	}
	for i := 0; i < len(endingArr); i++ {
		matchEndingArr = append(matchEndingArr, strArr[len(strArr)-1-i])
	}
	if strings.Join(matchEndingArr, "") == ending {
		return true
	}
	return false
}

// much easier way 
func betterSolution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}
