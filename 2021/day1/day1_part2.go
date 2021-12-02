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

	var numberOfIncreases int

	//Read first three elements
	sc.Scan()
	prevPrevPrevDepth, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	prevPrevDepth, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	prevDepth, _ := strconv.Atoi(sc.Text())

	for sc.Scan() {
		depth, _ := strconv.Atoi(sc.Text())
		if depth > prevPrevPrevDepth {
			numberOfIncreases++
		}
		prevPrevPrevDepth = prevPrevDepth
		prevPrevDepth = prevDepth
		prevDepth = depth
	}

	fmt.Println(numberOfIncreases)
}
