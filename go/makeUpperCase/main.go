package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MakeUpperCase("hello"))
	fmt.Println(MakeUpperCase("hello world"))
	fmt.Println(MakeUpperCase("hello world !"))
	fmt.Println(MakeUpperCase("heLlO wORLd !"))
	fmt.Println(MakeUpperCase("1,2,3 hello world!"))
}

func MakeUpperCase(str string) string {
	return strings.ToUpper(str)
}
