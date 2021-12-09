package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	draw := strings.Split(sc.Text(), ",")
	turn_draw := make(map[string]int)
	for turn, number := range draw {
		turn_draw[number] = turn
	}
	sc.Scan()

	var worst_board [][]string
	win_turn_of_worst := 0

	var curr_board [][]string
	for {
		sc.Scan()
		if sc.Text() == "" {
			if len(curr_board) == 0 {
				break
			}

			win_turn := len(turn_draw)

			var lines_last_turns, cols_last_turns [5]int
			for i := 0; i < len(curr_board); i++ {
				for j := 0; j < len(curr_board); j++ {
					cell_turn := turn_draw[curr_board[i][j]]
					if cell_turn > lines_last_turns[i] {
						lines_last_turns[i] = cell_turn
					}
					if cell_turn > cols_last_turns[j] {
						cols_last_turns[j] = cell_turn
					}
				}
			}
			for _, turn := range lines_last_turns {
				if turn < win_turn {
					win_turn = turn
				}
			}
			for _, turn := range cols_last_turns {
				if turn < win_turn {
					win_turn = turn
				}
			}
			if win_turn > win_turn_of_worst {
				worst_board = curr_board
				win_turn_of_worst = win_turn
			}
			curr_board = [][]string{}
		} else {
			curr_board = append(curr_board, strings.Fields(sc.Text()))
		}
	}

	var sum int
	for _, line := range worst_board {
		for _, number := range line {
			int_number, _ := strconv.Atoi(number)
			if turn_draw[number] > win_turn_of_worst {
				sum += int_number
			}
		}
	}

	win_numvber, _ := strconv.Atoi(draw[win_turn_of_worst])
	fmt.Println(sum * win_numvber)
}
