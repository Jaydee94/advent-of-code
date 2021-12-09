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

	var best_board [][]string
	winning_turn_best := len(turn_draw)

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
					cellTurn := turn_draw[curr_board[i][j]]
					if cellTurn > lines_last_turns[i] {
						lines_last_turns[i] = cellTurn
					}
					if cellTurn > cols_last_turns[j] {
						cols_last_turns[j] = cellTurn
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
			if win_turn < winning_turn_best {
				best_board = curr_board
				winning_turn_best = win_turn
			}
			curr_board = [][]string{}
		} else {
			curr_board = append(curr_board, strings.Fields(sc.Text()))
		}
	}

	var sum int
	for _, line := range best_board {
		for _, number := range line {
			intNumber, _ := strconv.Atoi(number)
			if turn_draw[number] > winning_turn_best {
				sum += intNumber
			}
		}
	}

	win_number, _ := strconv.Atoi(draw[winning_turn_best])
	fmt.Println(sum * win_number)
}
