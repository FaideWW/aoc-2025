package main

import (
	"fmt"
	"os"
	"strconv"

	io "github.com/faideww/aoc-2025/lib"
)

func main() {
	input := io.ReadInputFile(os.Args[1])
	lines := io.TrimAndSplit(input)

	count := countZeros(lines)
	fmt.Printf("count:%d\n", count)
}

func countZeros(lines []string) int {
	zeroes := 0
	current := 50

	for _, line := range lines {
		prev := current
		dir := line[0]
		dist, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}

		if dir == 'L' {
			startedAtZero := prev == 0
			current -= dist
			for current < 0 {
				if startedAtZero {
					zeroes--
					startedAtZero = false
				}
				zeroes++
				current += 100
			}

			if current == 0 {
				zeroes++
			}

		} else {
			current += dist
			for current > 99 {
				zeroes++
				current -= 100
			}
		}
	}

	return zeroes
}
