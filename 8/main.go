package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"

	io "github.com/faideww/aoc-2025/lib"
)

type Pos struct {
	x, y, z int
}

type Pair struct {
	p1         Pos
	p2         Pos
	p1Id, p2Id int
	dist       float64
}

type Circuit struct {
	points map[int]struct{}
}

func main() {
	input := io.ReadInputFileNoTrim(os.Args[1])
	lines := io.TrimAndSplit(input)

	points, pairs := findPairs(lines)

	// part 1
	result := connectBoxes(pairs, 1000)
	fmt.Printf("part1:%d\n", result)

	// part 2
	result2 := connectEverything(points, pairs)
	fmt.Printf("part2:%d\n", result2)
}

func findPairs(lines []string) ([]Pos, []Pair) {
	positions := make([]Pos, len(lines))
	for i, line := range lines {
		coords := strings.Split(line, ",")
		x, xerr := strconv.Atoi(coords[0])
		if xerr != nil {
			panic(xerr)
		}
		y, yerr := strconv.Atoi(coords[1])
		if yerr != nil {
			panic(yerr)
		}
		z, zerr := strconv.Atoi(coords[2])
		if zerr != nil {
			panic(zerr)
		}
		positions[i] = Pos{x, y, z}
	}

	sorted := sortClosestPairs(positions)
	return positions, sorted
}

func connectEverything(points []Pos, pairs []Pair) int {
	circuits := []Circuit{}
	circuitsByPoint := make(map[int]int)
	for i := range points {
		c := Circuit{
			points: make(map[int]struct{}),
		}
		c.points[i] = struct{}{}
		circuitsByPoint[i] = len(circuits)
		circuits = append(circuits, c)
	}

	circuitCount := len(circuits)

	nextPair := 0
	for circuitCount > 1 {
		pair := pairs[nextPair]

		c1Id := circuitsByPoint[pair.p1Id]
		c2Id := circuitsByPoint[pair.p2Id]

		if c1Id == c2Id {
			// If these are the same circuit, they're already merged
		} else {
			c1 := circuits[c1Id]
			c2 := circuits[c2Id]
			for pointId := range c2.points {
				c1.points[pointId] = struct{}{}
				circuitsByPoint[pointId] = c1Id
			}
			// circuits = append(circuits[:c2Id], circuits[c2Id+1:]...)
			circuitCount--
		}

		nextPair++
	}

	lastPair := pairs[nextPair-1]
	return lastPair.p1.x * lastPair.p2.x
}

func connectBoxes(pairs []Pair, connections int) int {
	circuits := []Circuit{}

	for i := range connections {
		pair := pairs[i]
		c := Circuit{
			points: make(map[int]struct{}),
		}
		c.points[pair.p1Id] = struct{}{}
		c.points[pair.p2Id] = struct{}{}
		circuits = append(circuits, c)
	}

	// Merge circuits
	merged := true
	for merged {
		merged = false
		for i := 0; i < len(circuits); i++ {
			c1 := circuits[i]
			for j := i + 1; j < len(circuits); j++ {
				c2 := circuits[j]

				for id := range c1.points {
					_, ok := c2.points[id]
					if ok {
						// Merge!
						merged = true
						for id2 := range c2.points {
							c1.points[id2] = struct{}{}
						}
						circuits = append(circuits[:j], circuits[j+1:]...)

						break
					}
				}

				if merged == true {
					break
				}
			}

			if merged == true {
				break
			}
		}
	}

	slices.SortFunc(circuits, func(a, b Circuit) int {
		return len(b.points) - len(a.points)
	})

	product := 1
	for i := range 3 {
		product *= len(circuits[i].points)
	}
	return product
}

func sortClosestPairs(points []Pos) []Pair {
	sorted := []Pair{}
	for i := range len(points) {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]
			dist := distance(p1, p2)
			sorted = append(sorted, Pair{p1, p2, i, j, dist})

		}
	}

	slices.SortFunc(sorted, func(a, b Pair) int {
		res := a.dist - b.dist
		if res < 0 {
			return -1
		} else {
			return 1
		}
	})
	return sorted
}

func distance(p1, p2 Pos) float64 {
	return math.Sqrt(
		math.Pow(float64(p2.x-p1.x), 2) +
			math.Pow(float64(p2.y-p1.y), 2) +
			math.Pow(float64(p2.z-p1.z), 2))
}
