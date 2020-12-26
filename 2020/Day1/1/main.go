package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	set map[int]bool
	num int
)

func main() {
	// init
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading input:\n", err)
		os.Exit(-1)
	}

	set = make(map[int]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, _ = strconv.Atoi(scanner.Text())
		// fmt.Println(num)
		set[num] = true
		if set[2020-num] {
			fmt.Println(num * (2020 - num))
			break
		}
	}

}
