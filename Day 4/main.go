package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("assign2.txt")

	if err != nil {
		fmt.Println(err)
	}

	temp := strings.Split(string(readFile), "\n")

	contains := 0
	dups := 0

	for _, line := range temp {
		test := strings.Split(line, ",")
		firstH := strings.Split(test[0], "-")
		secondH := strings.Split(test[1], "-")

		min1, _ := strconv.Atoi(firstH[0])
		max1, _ := strconv.Atoi(firstH[1])
		min2, _ := strconv.Atoi(secondH[0])
		max2, _ := strconv.Atoi(secondH[1])

		if (min1 >= min2 && max1 <= max2) || (min2 >= min1 && max2 <= max1) {
			contains += 1

		}
		if overLap(min1, max1, min2, max2) {
			fmt.Println(firstH, secondH)
			dups += 1
		}
	}
	fmt.Println("Fully contains:", contains)
	fmt.Println("Dup count:", dups)
}

func overLap(min1, max1, min2, max2 int) bool {
	if min1 >= min2 && min1 <= max2 {
		return true
	}
	if min2 >= min1 && min2 <= max1 {
		return true
	}
	return false
}
