package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	//data := "mjqjpqmgbljsphdztnvjfqwrcgsmlb"

	char := strings.Split(string(data), "")
	count := make(map[string]struct{})
	var cur []string
	i := 0
	marker := 14

	for len(count) < marker {
		count = make(map[string]struct{})
		cur, char = char[:marker], char[1:]

		for _, char := range cur {
			count[char] = struct{}{}
		}

		i++
	}
	fmt.Println(i + marker - 1)

}
