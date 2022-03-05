package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(solve("a"))                                                      //0
	fmt.Println(solve("bcd"))                                                    //9
	fmt.Println(solve("zea"))                                                    //26
	fmt.Println(solve("az"))                                                     //26
	fmt.Println(solve("baz"))                                                    //26
	fmt.Println(solve("aeiou"))                                                  //0
	fmt.Println(solve("uaoczei"))                                                //29
	fmt.Println(solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes")) //143
	fmt.Println(solve("codewars"))                                               //37
	fmt.Println(solve("bup"))                                                    //16
}

func solve(str string) int {
	consonants := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12, "m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18, "s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24, "y": 25, "z": 26}

	reg := regexp.MustCompile(`[^aeiou]+`)
	filteredArr := reg.FindAllStringSubmatch(str, -1)
	result := []int{}
	for _, subStr := range filteredArr {
		total := 0
		for _, eachSubStr := range subStr {
			for _, char := range strings.Split(eachSubStr, "") {
				total = total + consonants[char]
			}
			result = append(result, total)
			total = 0
		}
	}
	sort.Ints(result)
	if len(result) == 0 {
		return 0
	}
	return result[len(result)-1]
}
