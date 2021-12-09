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
	var fishes_by_waiting_time [9]int

	for _, fish := range strings.Split(sc.Text(), ",") {
		fish_waiting_time, _ := strconv.Atoi(fish)
		fishes_by_waiting_time[fish_waiting_time]++
	}

	for generation := 0; generation < 80; generation++ {
		just_bred := fishes_by_waiting_time[0]
		for days_wait := range fishes_by_waiting_time[:len(fishes_by_waiting_time)-1] {
			fishes_by_waiting_time[days_wait] = fishes_by_waiting_time[days_wait+1]
		}
		fishes_by_waiting_time[6] += just_bred
		fishes_by_waiting_time[8] = just_bred
	}

	var num_fishes int
	for _, fishes := range fishes_by_waiting_time {
		num_fishes += fishes
	}
	fmt.Println(num_fishes)
}
