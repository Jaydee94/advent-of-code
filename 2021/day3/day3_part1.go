package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var sum_columns [12]int
	var lines int

	for sc.Scan() {
		for i, bit := range sc.Text() {
			sum_columns[i] += int(bit - 48) // 48 ASCII Code of '0'
		}
		lines++
	}

	var gamma_rate, epsilon_rate int
	for i, sum := range sum_columns {
		if sum > lines/2 {
			gamma_rate += int(math.Pow(2, float64(len(sum_columns)-1-i)))
		} else {
			epsilon_rate += int(math.Pow(2, float64(len(sum_columns)-1-i)))
		}
	}
	fmt.Println(gamma_rate * epsilon_rate)
}
