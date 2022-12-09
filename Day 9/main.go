package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coor struct {
	x int
	y int
}

func main() {
	input, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(follows(string(input), 10))
}

func follows(str string, knots int) int {
	rope := []*coor{}

	for i := 0; i < knots; i++ {
		rope = append(rope, &coor{0, 0})
	}

	tailPos := map[string]int{"0 0": 1}

	d2m := map[coor]coor{}

	d2m[coor{0, -2}] = coor{0, -1}
	d2m[coor{0, 2}] = coor{0, 1}
	d2m[coor{-2, 0}] = coor{-1, 0}
	d2m[coor{2, 0}] = coor{1, 0}

	d2m[coor{2, 1}] = coor{1, 1}
	d2m[coor{2, -1}] = coor{1, -1}
	d2m[coor{-2, 1}] = coor{-1, 1}
	d2m[coor{-2, -1}] = coor{-1, -1}

	d2m[coor{1, 2}] = coor{1, 1}
	d2m[coor{1, -2}] = coor{1, -1}
	d2m[coor{-1, 2}] = coor{-1, 1}
	d2m[coor{-1, -2}] = coor{-1, -1}

	d2m[coor{-2, -2}] = coor{-1, -1}
	d2m[coor{2, 2}] = coor{1, 1}
	d2m[coor{-2, 2}] = coor{-1, 1}
	d2m[coor{2, -2}] = coor{1, -1}

	data := strings.Split(str, "\n")

	for _, v := range data {
		f := strings.Fields(v)
		direction := f[0]
		dist, _ := strconv.Atoi(f[1])
		moveQ := ""
		for m := 0; m < dist; m++ {
			moveQ += direction
		}

		for len(moveQ) > 0 {
			move := moveQ[0]
			moveQ = moveQ[1:]

			head := rope[0]

			switch move {
			case 'U':
				head.y += 1
			case 'D':
				head.y -= 1
			case 'L':
				head.x -= 1
			case 'R':
				head.x += 1
			}

			for k, v := range rope {
				if k == 0 {
					continue
				}

				move := d2m[coor{rope[k-1].x - v.x, rope[k-1].y - v.y}]
				v.x += move.x
				v.y += move.y
				if k == len(rope)-1 {
					tailPos[fmt.Sprintf("%d %d", v.x, v.y)] += 1
				}
			}
		}
	}

	return len(tailPos)
}
