package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	io "github.com/faideww/aoc-2025/lib"
)

type Point struct {
	x int
	y int
}

type Pair struct {
	p1, p2 Point
}

func main() {
	input := io.ReadInputFileNoTrim(os.Args[1])
	lines := io.TrimAndSplit(input)

	points := parsePoints(lines)
	allPairs := pairwisePoints(points)
	area := findLargestRect(allPairs)
	fmt.Printf("part1: %d\n", area)

	interiorPairs := filterPairs(allPairs, points)
	area2 := findLargestRect(interiorPairs)
	fmt.Printf("part2: %d\n", area2)
}

func parsePoints(lines []string) []Point {
	p := make([]Point, len(lines))
	for i, l := range lines {
		components := strings.Split(l, ",")
		x, xerr := strconv.Atoi(components[0])
		if xerr != nil {
			panic(xerr)
		}
		y, yerr := strconv.Atoi(components[1])
		if yerr != nil {
			panic(yerr)
		}
		p[i] = Point{x, y}

	}

	return p
}

func pairwisePoints(points []Point) []Pair {
	pairs := []Pair{}
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			pairs = append(pairs, Pair{p1, p2})
		}
	}
	return pairs
}

// Hypothesis: assuming we are dealing with a simple polygon
// (ie. does not have intersecting sides), a sufficient test
// should be that none of the polygon's points lie inside the rectangle,
// AND all of the rectangle's points lie inside the polygon.
func filterPairs(pairs []Pair, polygon []Point) []Pair {
	fmt.Printf("polygon:%v\n", polygon)
	fmt.Printf("filter 1\n")
	filteredPairs := []Pair{}
	for _, pair := range pairs {
		tl := Point{x: min(pair.p1.x, pair.p2.x), y: min(pair.p1.y, pair.p2.y)}
		br := Point{x: max(pair.p1.x, pair.p2.x), y: max(pair.p1.y, pair.p2.y)}

		validRect := true
		for i, pp1 := range polygon {
			j := i + 1
			if j == len(polygon) {
				j = 0
			}
			pp2 := polygon[j]

			var minX, maxX, minY, maxY int
			if pp1.x < pp2.x {
				minX = pp1.x
				maxX = pp2.x
			} else {
				minX = pp2.x
				maxX = pp1.x
			}
			if pp1.y < pp2.y {
				minY = pp1.y
				maxY = pp2.y
			} else {
				minY = pp2.y
				maxY = pp1.y
			}
			if tl.x < maxX && br.x > minX && tl.y < maxY && br.y > minY {
				// Do an aabb test to see if the edge intersects the rect
				validRect = false
				break
			}
		}
		if validRect {
			filteredPairs = append(filteredPairs, pair)
		}
	}

	return filteredPairs
}

func rectArea(p1, p2 Point) int {
	h := p2.y - p1.y
	if h < 0 {
		h = h * -1
	}
	h++

	w := p2.x - p1.x
	if w < 0 {
		w = w * -1
	}
	w++

	return w * h
}

func findLargestRect(pairs []Pair) int {
	maxArea := 0
	for _, pair := range pairs {
		p1 := pair.p1
		p2 := pair.p2
		area := rectArea(p1, p2)
		if maxArea < area {
			maxArea = area
		}
	}
	return maxArea
}
