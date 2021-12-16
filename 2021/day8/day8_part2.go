package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Read the file
	input, _ := os.Open("input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var sum int
	for sc.Scan() {
		signal := strings.Split(sc.Text(), "|")
		sig := decode(strings.Fields(signal[0]))

		decoded_display := ""
		for _, digit := range strings.Fields(signal[1]) {
			decoded_display += sig[sortSegments(digit)]
		}
		value, _ := strconv.Atoi(decoded_display)
		sum += value
	}
	fmt.Println(sum)
}

func decode(signal []string) map[string]string {
	sig := make(map[string]string)
	count_instance := make(map[rune]int)

	for _, digit := range signal {
		for _, segment := range digit {
			count_instance[segment]++
		}
	}

	for _, digit := range signal {
		switch len(digit) {
		case 2:
			sig[sortSegments(digit)] = "1"
		case 3:
			sig[sortSegments(digit)] = "7"
		case 4:
			sig[sortSegments(digit)] = "4"
		case 5:
			for _, segment := range digit {
				if count_instance[segment] == 4 {
					sig[sortSegments(digit)] = "2"
				} else if count_instance[segment] == 6 {
					sig[sortSegments(digit)] = "5"
				}
			}
			if _, twoOrFive := sig[sortSegments(digit)]; !twoOrFive {
				sig[sortSegments(digit)] = "3"
			}
		case 6:
			count_segments_with_key_instances := make(map[int]int)
			for _, segment := range digit {
				count_segments_with_key_instances[count_instance[segment]]++
			}
			if count_segments_with_key_instances[4] == 0 {
				sig[sortSegments(digit)] = "9"
			} else if count_segments_with_key_instances[7] == 1 {
				sig[sortSegments(digit)] = "0"
			} else {
				sig[sortSegments(digit)] = "6"
			}
		case 7:
			sig[sortSegments(digit)] = "8"
		}
	}
	return sig
}

func sortSegments(digit string) (sorted string) {
	for _, segment := range []string{"a", "b", "c", "d", "e", "f", "g"} {
		if strings.Contains(digit, segment) {
			sorted += segment
		}
	}
	return
}
