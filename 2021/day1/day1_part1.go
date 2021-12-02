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

	var numberOfIncreases int

	sc.Scan()
	prevDepth, _ := strconv.Atoi(sc.Text())
	for sc.Scan() {
		depth, _ := strconv.Atoi(sc.Text())
		if depth > prevDepth {
			numberOfIncreases++
		}
		prevDepth = depth
	}

	fmt.Println(numberOfIncreases)
}
