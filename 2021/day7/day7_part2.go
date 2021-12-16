package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	values := strings.Split(sc.Text(), ",")

	mean := 0.0
	for _, position := range values {
		pos_value, _ := strconv.Atoi(position)
		mean += float64(pos_value)
	}
	mean /= float64(len(values))

	align := int(mean)

	cost := 0
	for _, position := range values {
		pos_value, _ := strconv.Atoi(position)
		cost += fuel(int(math.Abs(float64(pos_value - align))))
	}
	fmt.Println(cost)
}

func fuel(n int) int {
	if n == 0 {
		return 0
	}
	return n + fuel(n-1)
}
