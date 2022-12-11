package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	//const screenWidth = 40
	var cycle = 1
	var answer int
	var register = 1
	var line string
	screenWidth := 40
	var build strings.Builder

	data, _ := os.ReadFile("data.txt")
	lines := strings.Split(string(data), "\r\n")

	for len(lines) > 0 {
		if (cycle-20)%40 == 0 {
			answer += cycle * register
		}

		pixel := (cycle - 1) % screenWidth
		if pixel == 0 {
			build.WriteString("\n")
		}
		if pixel-1 <= register && register <= pixel+1 {
			build.WriteString("#")
		} else {
			build.WriteString(" ")
		}

		if len(line) == 0 {
			line = lines[0]
			if line == "noop" {
				lines = lines[1:]
				line = ""
			}

			cycle++
		} else {
			parsed, _ := strconv.Atoi(strings.Fields(line)[1])
			register += parsed
			lines = lines[1:]
			line = ""
			cycle++
		}
	}

	fmt.Println(build.String())
}
