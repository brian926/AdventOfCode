package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	readFile, err := os.ReadFile("rucksack.txt")

	if err != nil {
		fmt.Println(err)
	}

	temp := strings.Split(string(readFile), "\n")
	total := everyThree(temp)

	fmt.Println(total)
}

func everyThree(temp []string) int {
	ran := 3
	total := 0
	for i := 0; ran <= len(temp); i += 3 {
		test := temp[i:ran]
		fmt.Println(i, ran)
		letter := overLap(test[0], test[1], test[2])
		total += value(int(letter))
		ran += 3
	}
	return total
}

func value(common int) int {
	total := 0

	if common >= 97 {
		value := int(common) - 96
		total += value
	} else {
		value := int(common) - 38
		total += value
	}

	return total
}

func overLap(firstH, secondH, thirdH string) rune {
	var ret rune
	for _, l := range firstH {
		if isMember(secondH, l) && isMember(thirdH, l) {
			ret = l
			break
		}
	}
	return ret
}

func isMember(secondH string, letter rune) bool {
	for _, l := range secondH {
		if letter == l {
			return true
		}
	}
	return false
}
