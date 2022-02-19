package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(ToAlternatingCase("hello world"))
	fmt.Println(ToAlternatingCase("HELLO WORLD"))
	fmt.Println(ToAlternatingCase("hello WORLD"))
	fmt.Println(ToAlternatingCase("HeLLo WoRLD"))
	fmt.Println(ToAlternatingCase("12345"))
	fmt.Println(ToAlternatingCase("1a2b3c4d5e"))
}

func ToAlternatingCase(str string) string {
	ans := []rune{}
	for _, rune := range str {
		if unicode.IsUpper(rune) {
			ans = append(ans, unicode.ToLower(rune))
		} else {
			ans = append(ans, unicode.ToUpper(rune))
		}
	}
	return string(ans)
}
