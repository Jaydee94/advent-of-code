package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var unique_patterns [][]string
	var four_digit [][]string

	for sc.Scan() {
		signal := strings.Split(sc.Text(), "|")
		unique_patterns = append(unique_patterns, strings.Fields(signal[0]))
		four_digit = append(four_digit, strings.Fields(signal[1]))
	}

	var istances int

	for _, display := range four_digit {
		for _, digit := range display {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				istances++
			}
		}
	}

	fmt.Println(istances)
}
