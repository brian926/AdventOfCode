package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.ReadFile("calories.txt")

	if err != nil {
		fmt.Println(err)
	}

	temp := strings.Split(string(readFile), "\n")

	var arr []int
	count := 0

	for _, line := range temp {
		if line != "" {
			car, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}

			count = count + car
		} else {
			arr = append(arr, count)
			count = 0
		}

	}

	sort.Ints(arr)
	lenA := len(arr)
	max := arr[lenA-1] + arr[lenA-2] + arr[lenA-3]

	fmt.Println(max)
}
