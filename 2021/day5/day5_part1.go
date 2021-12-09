package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	lines_in_point := make(map[string]int)

	for sc.Scan() {
		var start, end [2]int
		fmt.Sscanf(sc.Text(), "%d,%d -> %d,%d", &(start[0]), &(start[1]), &(end[0]), &(end[1]))

		if start[0] == end[0] || start[1] == end[1] {
			lines_in_point[fmt.Sprintf("%d,%d", start[0], start[1])]++
			for start[0] != end[0] || start[1] != end[1] {
				//We go ahead by one point near to the end
				if start[0] > end[0] {
					start[0]--
				} else if start[0] < end[0] {
					start[0]++
				}
				if start[1] > end[1] {
					start[1]--
				} else if start[1] < end[1] {
					start[1]++
				}
				lines_in_point[fmt.Sprintf("%d,%d", start[0], start[1])]++
			}
		}
	}

	var count_points int
	for _, v := range lines_in_point {
		if v > 1 {
			count_points++
		}
	}
	fmt.Println(count_points)
}
