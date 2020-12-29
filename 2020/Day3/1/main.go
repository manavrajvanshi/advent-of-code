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

	// for s := range terrain {
	// 	fmt.Println(terrain[s])
	// }

	i, j := 0, 0
	w, h := len(terrain[0]), len(terrain)
	count := 0
	for i < h {
		if string(terrain[i][j]) == "#" {
			count++
		}
		i++
		j = (j + 3) % w
	}
	fmt.Println(count)

}
