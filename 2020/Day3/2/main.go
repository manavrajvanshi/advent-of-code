package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// init
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading input:\n", err)
		os.Exit(-1)
	}

	var terrain []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		terrain = append(terrain, scanner.Text())
	}

	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	multiplied := 1
	for _, slope := range slopes {
		// fmt.Println(slope)
		i, j := 0, 0
		w, h := len(terrain[0]), len(terrain)
		count := 0
		for i < h {
			if string(terrain[i][j]) == "#" {
				count++
			}
			i = i + slope[1]
			j = (j + slope[0]) % w
		}
		multiplied *= count
	}
	fmt.Println(multiplied)

}
