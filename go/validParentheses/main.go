package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(ValidParentheses("()"))             // true
	fmt.Println(ValidParentheses(")"))              // false
	fmt.Println(ValidParentheses("(())((()())())")) // true
	fmt.Println(ValidParentheses("hi())("))         // false
	fmt.Println(ValidParentheses("hi(hi)"))         // true
	fmt.Println(ValidParentheses("())(()"))         // false
}

func ValidParentheses(parens string) bool {
	r := regexp.MustCompile(`[()]`)
	parenArr := []string{}
	strip := r.FindAllString(parens, -1)
	for i, bracket := range strip {
		if i == 0 {
			parenArr = append(parenArr, bracket)
		} else {
			if parenArr[i-1] == ")" {
				parenArr = parenArr[:i-1]
			} else {
				parenArr = append(parenArr, bracket)
			}
		}
	}
	fmt.Println(parenArr)
	return false
}
