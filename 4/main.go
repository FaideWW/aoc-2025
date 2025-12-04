package main

import (
	"fmt"
	"os"

	io "github.com/faideww/aoc-2025/lib"
)

type Pos struct {
	x int
	y int
}

func main() {
	input := io.ReadInputFile(os.Args[1])
	lines := io.TrimAndSplit(input)

	count := countAccessibleRolls(lines)
	fmt.Printf("accessible:%d\n", count)

	count2 := recursiveCount(lines)
	fmt.Printf("recursive:%d\n", count2)
}

func countAccessibleRolls(lines []string) int {
	count := 0
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] == '@' && isRollAccessible(lines, x, y) {
				count++
			}
		}
	}

	return count
}

func recursiveCount(lines []string) int {
	count := 0
	for {
		iterationCount := 0
		accessibles := make(map[Pos]struct{})
		for y := 0; y < len(lines); y++ {
			for x := 0; x < len(lines[y]); x++ {
				if lines[y][x] == '@' && isRollAccessible(lines, x, y) {
					count++
					iterationCount++
					accessibles[Pos{x, y}] = struct{}{}
				}
			}
		}

		if iterationCount == 0 {
			break
		}

		for p := range accessibles {
			lines[p.y] = fmt.Sprintf("%s%s%s", lines[p.y][:p.x], ".", lines[p.y][p.x+1:])
		}

		iterationCount = 0
	}
	return count
}

func isRollAccessible(lines []string, x, y int) bool {
	rollNeighbors := 0

	// NW
	if y > 0 && x > 0 && lines[y-1][x-1] == '@' {
		rollNeighbors++
	}

	// N
	if y > 0 && lines[y-1][x] == '@' {
		rollNeighbors++
	}

	// NE
	if y > 0 && x < len(lines[y-1])-1 && lines[y-1][x+1] == '@' {
		rollNeighbors++
	}

	// W
	if x > 0 && lines[y][x-1] == '@' {
		rollNeighbors++
	}

	// E
	if x < len(lines[y])-1 && lines[y][x+1] == '@' {
		rollNeighbors++
	}

	// SW
	if y < len(lines)-1 && x > 0 && lines[y+1][x-1] == '@' {
		rollNeighbors++
	}

	// S
	if y < len(lines)-1 && lines[y+1][x] == '@' {
		rollNeighbors++
	}

	// SE
	if y < len(lines)-1 && x < len(lines[y+1])-1 && lines[y+1][x+1] == '@' {
		rollNeighbors++
	}

	return rollNeighbors < 4
}
