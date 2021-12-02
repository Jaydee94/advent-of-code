package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	h_pos := 0
	depth := 0
	aim := 0

	for _, line := range txtlines {
		res := strings.Split(line, " ")
		command := res[0]
		counter, _ := strconv.Atoi(res[1])
		switch {
		case command == "forward":
			h_pos = h_pos + counter
			depth = depth + (aim * counter)
		case command == "down":
			aim = aim + counter
		default:
			aim = aim - counter
		}
	}
	fmt.Println("horizontal_pos: ", h_pos)
	fmt.Println("depth: ", depth)
	fmt.Println("result: ", h_pos*depth)
}
