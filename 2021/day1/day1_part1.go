package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var increases int

	sc.Scan()
	prev_depth, _ := strconv.Atoi(sc.Text())
	for sc.Scan() {
		depth, _ := strconv.Atoi(sc.Text())
		if depth > prev_depth {
			increases++
		}
		prev_depth = depth
	}

	fmt.Println(increases)
}
