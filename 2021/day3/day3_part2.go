package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var oxygen_rating []string
	for sc.Scan() {
		oxygen_rating = append(oxygen_rating, sc.Text())
	}
	co2_rating := append([]string{}, oxygen_rating...)

	gamma, _ := strconv.ParseInt(filer_values(oxygen_rating, 0, true), 2, 64)
	epsilon, _ := strconv.ParseInt(filer_values(co2_rating, 0, false), 2, 64)
	fmt.Println(gamma * epsilon)
}

func filer_values(values []string, bit_to_consider int, most_common bool) string {
	if len(values) == 1 {
		return values[0]
	}
	var value_zero, value_one []string

	for _, value := range values {
		if rune(value[bit_to_consider]) == '0' {
			value_zero = append(value_zero, value)
		} else {
			value_one = append(value_one, value)
		}
	}
	if len(value_one) >= len(value_zero) == most_common {
		return filer_values(value_one, bit_to_consider+1, most_common)
	}
	return filer_values(value_zero, bit_to_consider+1, most_common)
}
