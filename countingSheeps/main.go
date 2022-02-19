package main

func main() {

}

func CountSheeps(numbers []bool) int {
	num := 0
	for _, sheep := range numbers {
		if sheep {
			num++
		}
	}
	return num
}
