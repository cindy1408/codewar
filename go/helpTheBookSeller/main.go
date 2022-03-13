package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	fmt.Println(StockList(b, c))
}

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	organise := make(map[string]int)
	for _, char := range listCat {
		organise[char] = 0
	}
	for _, eachBook := range listArt {
		listArtArr := strings.Split(eachBook, " ")
		for _, char := range listCat {
			if strings.HasPrefix(listArtArr[0], char) {
				num, _ := strconv.Atoi(listArtArr[1])
				organise[char] = organise[char] + num
			}
		}
	}
	return format(organise, listCat)
}

func format(organise map[string]int, listCat []string) string {
	formatArr := []string{}
	for _, char := range listCat {
		for index, sum := range organise {
			if char == index {
				formatArr = append(formatArr, fmt.Sprintf("(%s : %s)", char, strconv.Itoa(sum)))
			}
		}
	}
	return strings.Join(formatArr, " - ")
}
