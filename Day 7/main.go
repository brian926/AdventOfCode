package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	value, err := os.ReadFile("fileS.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := strings.Split(string(value), "\n")

	sizes := map[string]int{"/": 0}
	dirPath := []string{"/"}

	re := regexp.MustCompile("[0-9]+")

	for _, cmd := range data {
		if cmd == "$ cd /" {
			dirPath = []string{"/"}
			continue
		}
		if cmd == "$ cd .." {
			dirPath = dirPath[:len(dirPath)-1]
			continue
		}
		if cmd[:4] == "$ cd" {
			path := strings.Split(cmd, " ")[2]
			dirPath = append(dirPath, path)
			continue
		}
		if cmd[:3] == "dir" || cmd == "$ ls" {
			continue
		}

		num, _ := strconv.Atoi(re.FindAllString(cmd, 1)[0])

		currDirPath := "/"
		sizes[currDirPath] += num

		for _, dir := range dirPath[1:] {
			currDirPath += string(dir) + "/"
			sizes[currDirPath] += num
		}
	}

	sum := 0
	for _, size := range sizes {
		if size <= 100000 {
			sum += size
		}
	}

	minSize := sizes["/"] - 40000000
	currentMin := sizes["/"]
	for _, size := range sizes {
		if size > minSize && size < currentMin {
			currentMin = size
		}
	}
	fmt.Println(currentMin)
}
