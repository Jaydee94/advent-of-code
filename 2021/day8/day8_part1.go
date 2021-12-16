package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var unique [][]string
	var four_digit [][]string

	for sc.Scan() {
		signal := strings.Split(sc.Text(), "|")
		unique = append(unique, strings.Fields(signal[0]))
		four_digit = append(four_digit, strings.Fields(signal[1]))
	}

	var instance int

	for _, display := range four_digit {
		for _, digit := range display {
			if len(digit) == 2 || len(digit) == 3 || len(digit) == 4 || len(digit) == 7 {
				instance++
			}
		}
	}

	fmt.Println(instance)
}
