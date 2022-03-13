package main

import "strings"

func main() {

}

func RepeatStr(repetitions int, value string) string {
	repetition := []string{}
	for i := 0; i < repetitions; i++ {
		repetition = append(repetition, value)
	}
	return strings.Join(repetition, "")
}
