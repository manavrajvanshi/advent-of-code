package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// init
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading input:\n", err)
		os.Exit(-1)
	}
	scanner := bufio.NewScanner(file)
	var arrSlice []int
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		arrSlice = append(arrSlice, num)
	}

	sort.Ints(arrSlice)
	for i := 0; i < len(arrSlice)-2; i++ {
		l := i + 1
		r := len(arrSlice) - 1

		for l < r {
			sum := arrSlice[i] + arrSlice[l] + arrSlice[r]
			if sum == 2020 {
				fmt.Println(arrSlice[i] * arrSlice[l] * arrSlice[r])
				break
			} else if sum > 2020 {
				r--
			} else {
				l++
			}
		}
	}
}
