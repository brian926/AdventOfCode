package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	//str := strings.Split(string(data), "\n")

	mapData := parseMap(string(data))
	fmt.Println(calculateHighestScenicScore(mapData))
}

func parseMap(mapText string) [][]int {
	rowCount, columnCount := readDimensions(mapText)
	m := make([][]int, columnCount)

	for i := 0; i < len(m); i++ {
		m[i] = make([]int, rowCount)
	}

	readMap(mapText, m)
	return m
}

func readDimensions(mapText string) (rowCount, columnCount int) {
	columnCount = strings.Index(mapText, "\r\n")
	rowCount = strings.Count(mapText, "\r\n")

	if mapText[len(mapText)-1] != '\n' || mapText[len(mapText)-1] != '\r' {
		rowCount++
	}

	return
}

func readMap(mapText string, mapData [][]int) {
	lines := strings.Split(mapText, "\r\n")
	for i, line := range lines {
		for j, r := range line {
			x, _ := strconv.Atoi(string(r))
			mapData[i][j] = x
		}
	}
}

func calculateVisibleTrees(mapData [][]int) int {
	rowCount, columnCount := len(mapData[0]), len(mapData)
	visMap := make([][]bool, columnCount)

	for i := 0; i < len(visMap); i++ {
		visMap[i] = make([]bool, len(mapData[0]))
	}

	for i := 0; i < rowCount; i++ {
		if i == 0 || i == rowCount-1 {
			for j := 0; j < columnCount; j++ {
				visMap[i][j] = true
			}
		} else {
			visMap[i][0] = true
			visMap[i][columnCount-1] = true
		}
	}

	for i := 1; i < rowCount-1; i++ {
		for j := 1; j < columnCount-1; j++ {
			currentTree := mapData[i][j]
			blockingSides := 0

			for x := 0; x < j; x++ {
				if mapData[i][x] >= currentTree {
					blockingSides++
					break
				}
			}

			for y := j + 1; y < columnCount; y++ {
				if mapData[i][y] >= currentTree {
					blockingSides++
					break
				}
			}

			for u := 0; u < i; u++ {
				if mapData[u][j] >= currentTree {
					blockingSides++
					break
				}
			}

			for t := i + 1; t < rowCount; t++ {
				if mapData[t][j] >= currentTree {
					blockingSides++
					break
				}
			}

			visMap[i][j] = blockingSides < 4
		}
	}

	count := 0
	for i := 0; i < rowCount; i++ {
		for j := 0; j < columnCount; j++ {
			if visMap[i][j] {
				count++
			}
		}
	}

	return count
}

func calculateHighestScenicScore(mapData [][]int) int {
	rowCount, columnCount := len(mapData[0]), len(mapData)

	maxScore := 0

	for i := 1; i < rowCount-1; i++ {
		for j := 1; j < columnCount-1; j++ {
			currentTree := mapData[i][j]
			score := 1

			// x-axis search [leftside]
			for J := j - 1; 0 <= J; J-- {
				if mapData[i][J] >= currentTree || J == 0 {
					distance := j - J
					score *= distance
					break
				}
			}

			// x-axis search [rightside]
			for J := j + 1; J < columnCount; J++ {
				if mapData[i][J] >= currentTree || J == columnCount-1 {
					distance := J - j
					score *= distance
					break
				}
			}

			// y-axis search [topside]
			for I := i - 1; 0 <= I; I-- {
				if mapData[I][j] >= currentTree || I == 0 {
					distance := i - I
					score *= distance
					break
				}
			}

			// y-axis search [bottomside]
			for I := i + 1; I < rowCount; I++ {
				if mapData[I][j] >= currentTree || I == rowCount-1 {
					distance := I - i
					score *= distance
					break
				}
			}

			maxScore = max(maxScore, score)
		}
	}

	return maxScore
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
