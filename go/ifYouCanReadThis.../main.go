package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToNato("If you can read"))
	fmt.Println(ToNato("Did not see that coming"))
	fmt.Println(ToNato("go for it!"))
}

func ToNato(words string) string {
	result := []string{}
	for _, char := range strings.Split(words, "") {
		char = strings.ToLower(char)
		switch char {
		case "a":
			result = append(result, "Alfa")
		case "b":
			result = append(result, "Bravo")
		case "c":
			result = append(result, "Charlie")
		case "d":
			result = append(result, "Delta")
		case "e":
			result = append(result, "Echo")
		case "f":
			result = append(result, "Foxtrot")
		case "g":
			result = append(result, "Golf")
		case "h":
			result = append(result, "Hotel")
		case "i":
			result = append(result, "India")
		case "j":
			result = append(result, "Juliett")
		case "k":
			result = append(result, "Kilo")
		case "l":
			result = append(result, "Lima")
		case "m":
			result = append(result, "Mike")
		case "n":
			result = append(result, "November")
		case "o":
			result = append(result, "Oscar")
		case "p":
			result = append(result, "Papa")
		case "q":
			result = append(result, "Quebec")
		case "r":
			result = append(result, "Romeo")
		case "s":
			result = append(result, "Sierra")
		case "t":
			result = append(result, "Tango")
		case "u":
			result = append(result, "Uniform")
		case "v":
			result = append(result, "Victor")	
		case "w":
			result = append(result, "Whiskey")
		case "x":
			result = append(result, "X-ray")
		case "y":
			result = append(result, "Yankee")	
		case "z":
			result = append(result, "Zulu")	
		case ",":
			result = append(result, ",")
		case ".":
			result = append(result, ".")	
		case "!":
			result = append(result, "!")	
		case "?":
			result = append(result, "?")	
		}
	}
	return strings.Join(result, " ")
}