package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	scores := make(map[string]int)
	scores["lost"] = 0
	scores["draw"] = 3
	scores["won"] = 6

	shape := make(map[string]int)
	shape["X"] = 1
	shape["Y"] = 2
	shape["Z"] = 3

	won := make(map[string]string)
	won["A"] = "Y"
	won["B"] = "Z"
	won["C"] = "X"

	draw := make(map[string]string)
	draw["A"] = "X"
	draw["B"] = "Y"
	draw["C"] = "Z"

	lost := make(map[string]string)
	lost["A"] = "Z"
	lost["B"] = "X"
	lost["C"] = "Y"

	end := make(map[string]string)
	end["X"] = "lost"
	end["Y"] = "draw"
	end["Z"] = "won"

	readFile, err := os.ReadFile("strat.txt")

	if err != nil {
		fmt.Println(err)
	}

	total := 0

	temp := strings.Split(string(readFile), "\n")
	for _, line := range temp {
		value := strings.Split(line, " ")

		result := end[value[1]]

		var myShape string
		if result == "lost" {
			myShape = lost[value[0]]
		} else if result == "won" {
			myShape = won[value[0]]
		} else {
			myShape = draw[value[0]]
		}

		total += scores[result] + shape[myShape]

	}
	fmt.Println(total)
}
