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

	var increases int

	//Read first three elements
	sc.Scan()
	pre_pre_pre_depth, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	pre_pre_depth, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	pre_depth, _ := strconv.Atoi(sc.Text())

	for sc.Scan() {
		depth, _ := strconv.Atoi(sc.Text())
		if depth > pre_pre_pre_depth {
			increases++
		}
		pre_pre_pre_depth = pre_pre_depth
		pre_pre_depth = pre_depth
		pre_depth = depth
	}

	fmt.Println(increases)
}
