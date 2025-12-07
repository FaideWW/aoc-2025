package main

import (
	"fmt"
	"os"

	io "github.com/faideww/aoc-2025/lib"
)

type Pos struct {
	x, y int
}

func main() {
	input := io.ReadInputFileNoTrim(os.Args[1])
	lines := io.TrimAndSplit(input)

	count := countSplits(lines)
	fmt.Printf("splits:%d\n", count)

	timelines := countAllPaths(lines)
	fmt.Printf("timelines:%d\n", timelines)
}

func countSplits(lines []string) int {
	splits := 0
	width := len(lines[0])
	height := len(lines)

	beams := make(map[Pos]struct{})

	var start Pos

	for x := range width {
		if lines[0][x] == 'S' {
			start = Pos{x: x, y: 0}
		}
	}

	frontier := []Pos{start}
	beams[start] = struct{}{}

	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		if current.y < height-1 {
			y := current.y + 1
			x := current.x
			nextTile := lines[y][x]
			nextPos := Pos{x, y}

			if nextTile == '.' {
				if _, ok := beams[nextPos]; !ok {
					frontier = append(frontier, nextPos)
					beams[nextPos] = struct{}{}
				}
			} else if nextTile == '^' {
				splits++
				left := Pos{x: x - 1, y: y}
				right := Pos{x: x + 1, y: y}

				if _, ok := beams[left]; !ok {
					frontier = append(frontier, left)
					beams[left] = struct{}{}
				}

				if _, ok := beams[right]; !ok {
					frontier = append(frontier, right)
					beams[right] = struct{}{}
				}
			}
		}
	}

	return splits
}

func countAllPaths(lines []string) int {
	width := len(lines[0])
	height := len(lines)

	var start Pos

	for x := range width {
		if lines[0][x] == 'S' {
			start = Pos{x: x, y: 0}
		}
	}

	cachedTimelines := make(map[Pos]int)

	var countPaths func(current Pos) int
	countPaths = func(current Pos) int {
		next := current

		for next.y < height-1 && lines[next.y][next.x] != '^' {
			next.y++
		}
		if next.y < height-1 {
			nextTile := lines[next.y][next.x]

			if nextTile == '.' {
				panic(fmt.Errorf("somehow tile at %d,%d is not splitter", next.y, next.x))
			} else if nextTile == '^' {
				left := Pos{x: next.x - 1, y: next.y}
				right := Pos{x: next.x + 1, y: next.y}

				leftRoutes, leftOk := cachedTimelines[left]
				if !leftOk {
					leftRoutes = countPaths(left)
					cachedTimelines[left] = leftRoutes
				}

				rightRoutes, rightOk := cachedTimelines[right]
				if !rightOk {
					rightRoutes = countPaths(right)
					cachedTimelines[right] = rightRoutes
				}

				return leftRoutes + rightRoutes
			}
		}

		return 1

	}

	return countPaths(start)
}
