package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	s "strings"
)

func main() {
	// init
	file, err := os.Open("../input.txt")
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading input:\n", err)
		os.Exit(-1)
	}
	valid := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inp := s.Split(scanner.Text(), ": ")
		policy := s.Split(inp[0], " ")
		bounds := s.Split(policy[0], "-")

		l, _ := strconv.Atoi(bounds[0])
		h, _ := strconv.Atoi(bounds[1])
		str := inp[1]

		if isValid(l, h, policy[1], str) {
			valid++
		}
	}

	fmt.Println(valid)

}

func isValid(low, high int, char, str string) bool {
	runeStr := []rune(str)
	count := 0
	for i := 0; i < len(runeStr); i++ {
		if string(runeStr[i]) == char {
			count++
			if count > high {
				return false
			}
		}
	}
	return count >= low
}
