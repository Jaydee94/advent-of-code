package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var mat [][]rune

	for sc.Scan() {
		line := []rune{'9'}
		for _, point := range sc.Text() {
			line = append(line, point)
		}
		line = append(line, '9')
		mat = append(mat, line)
	}
	nines := make([]rune, len(mat[0]))
	for i := range nines {
		nines[i] = '9'
	}
	mat = append([][]rune{nines}, mat...)
	mat = append(mat, nines)

	low_points := make(map[string]bool)

	for i, line := range mat {
		for j, point := range line {
			if i != 0 && j != 0 &&
				i != len(mat)-1 && j != len(line)-1 &&
				point < mat[i+1][j] &&
				point < mat[i-1][j] &&
				point < mat[i][j+1] &&
				point < mat[i][j-1] {
				low_points[fmt.Sprintf("%d,%d", i, j)] = true
			}
		}
	}

	var max, max2, max3 int

	for point := range low_points {
		mat_cpy := append([][]rune{}, mat...)
		var x, y, b_points int
		fmt.Sscanf(point, "%d,%d", &x, &y)
		_, b_points = explore(mat_cpy, b_points, x, y)

		if b_points > max {
			max3 = max2
			max2 = max
			max = b_points
		} else if b_points > max2 {
			max3 = max2
			max2 = b_points
		} else if b_points > max3 {
			max3 = b_points
		}
	}
	fmt.Println(max * max2 * max3)
}

func explore(mat [][]rune, count, x, y int) ([][]rune, int) {
	if mat[x][y] != '9' {
		count++
		mat[x][y] = '9'
		mat, count = explore(mat, count, x+1, y)
		mat, count = explore(mat, count, x-1, y)
		mat, count = explore(mat, count, x, y+1)
		mat, count = explore(mat, count, x, y-1)
	}
	return mat, count
}
