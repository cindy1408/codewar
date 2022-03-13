package main

import (
	"fmt"
	"strings"
)

func main() {	
	fmt.Println(AbbrevName("Sam Harris"))
	fmt.Println(AbbrevName("Patrick Feenan"))
	fmt.Println(AbbrevName("Evan Cole"))
	fmt.Println(AbbrevName("P Favuzzi"))
}

func AbbrevName(name string) string {
	initals := []string{}
	nameArr := strings.Split(name, " ")
	for _, name := range nameArr {
		arr := strings.Split(name, "")
		initals = append(initals, strings.ToUpper(arr[0]))
	}
	return strings.Join(initals, ".")
}
