package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	d, err := os.ReadFile("crates.txt")
	if err != nil {
		fmt.Println(err)
	}

	data := strings.Split(string(d), "\n")

	amtBoxes := 9
	fmt.Println(data[0])

	boxes := make([][]string, amtBoxes)

	for _, v := range data[:amtBoxes] {
		for i := 0; i < amtBoxes; i++ {
			j := 4*i + 1
			if string(v[j]) == " " {
				continue
			}

			boxes[i] = append(boxes[i], string(v[j]))
		}
	}

	re := regexp.MustCompile("[0-9]+")

	for _, v := range data[amtBoxes+1:] {
		vals := re.FindAllString(v, 3)
		howMany, _ := strconv.Atoi(vals[0])
		from, _ := strconv.Atoi(vals[1])
		to, _ := strconv.Atoi(vals[2])

		from, to = from-1, to-1

		// for i := 0; i < howMany; i++ {
		// 	item := boxes[from][0]
		// 	boxes[from] = boxes[from][1:]
		// 	boxes[to] = append([]string{item}, boxes[to]...)
		// }

		items := make([]string, howMany)
		copy(items, boxes[from][0:howMany])
		boxes[from] = boxes[from][howMany:]
		boxes[to] = append(items, boxes[to]...)
	}

	ans := ""
	for _, v := range boxes {
		ans += v[0]
	}

	fmt.Println(ans)
}
