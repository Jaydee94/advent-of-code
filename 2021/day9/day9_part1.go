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

	risk := 0

	for i, line := range mat {
		for j, point := range line {
			if i != 0 && j != 0 &&
				i != len(mat)-1 && j != len(line)-1 &&
				point < mat[i+1][j] &&
				point < mat[i-1][j] &&
				point < mat[i][j+1] &&
				point < mat[i][j-1] {
				risk += int(point) - 48 + 1
			}
		}
	}
	fmt.Println(risk)
}
