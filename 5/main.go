package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	io "github.com/faideww/aoc-2025/lib"
)

type Range struct {
	low int64
	hi  int64
}

func main() {
	input := io.ReadInputFile(os.Args[1])

	components := io.TrimAndSplitBy(input, "\n\n")

	rangeLines := io.TrimAndSplit(components[0])
	ingredientLines := io.TrimAndSplit(components[1])

	ranges := parseRanges(rangeLines)
	ingredients := parseIngredients(ingredientLines)

	count := countFreshIngredients(ranges, ingredients)
	fmt.Printf("count:%d\n", count)

	count2 := countFreshRanges(ranges, ingredients)
	fmt.Printf("all:%d\n", count2)
}

func parseRanges(lines []string) []Range {
	r := []Range{}

	for _, line := range lines {
		vals := strings.Split(line, "-")
		low, err := strconv.ParseInt(vals[0], 10, 64)
		if err != nil {
			panic(err)
		}
		hi, err := strconv.ParseInt(vals[1], 10, 64)
		if err != nil {
			panic(err)
		}

		r = append(r, Range{low, hi})
	}

	return r
}

func parseIngredients(lines []string) []int64 {
	r := []int64{}

	for _, line := range lines {
		val, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}
		r = append(r, val)
	}

	return r
}

func countFreshIngredients(ranges []Range, ingredients []int64) int {
	total := 0
	for _, id := range ingredients {
		for _, r := range ranges {
			if id >= r.low && id <= r.hi {
				total++
				break
			}
		}
	}

	return total
}

func countFreshRanges(ranges []Range, ingredients []int64) int64 {
	dedupedRanges := ranges

	hasDupes := true
	for hasDupes {
		var i1, i2 int
		var newRange Range
		hasDupes, i1, i2, newRange = findRangeOverlap(dedupedRanges)
		if hasDupes {
			newDedupedRanges := append(dedupedRanges[:i1], dedupedRanges[i1+1:i2]...)
			newDedupedRanges = append(newDedupedRanges, dedupedRanges[i2+1:]...)
			newDedupedRanges = append(newDedupedRanges, newRange)
			dedupedRanges = newDedupedRanges
		}
	}

	count := int64(0)
	for _, r := range dedupedRanges {
		count += r.hi - r.low + 1
	}
	return count
}

func findRangeOverlap(ranges []Range) (hasOverlap bool, i1 int, i2 int, newRange Range) {
	hasOverlap = false
	for i := 0; i < len(ranges); i++ {
		r1 := ranges[i]
		i1 = i
		for j := i + 1; j < len(ranges); j++ {
			r2 := ranges[j]
			i2 = j

			switch {
			// case 1: [----r1--{-]--r2----}
			case r1.low <= r2.low && r1.low <= r2.hi && r1.hi >= r2.low && r1.hi <= r2.hi:
				{
					hasOverlap = true
					newRange = Range{r1.low, r2.hi}
					return
				}
			// case 2:          {----r2--[-}--r1----]
			case r1.low >= r2.low && r1.low <= r2.hi && r1.hi >= r2.low && r1.hi >= r2.hi:
				{
					hasOverlap = true
					newRange = Range{r2.low, r1.hi}
					return
				}
			// case 3: [--{-r2-}--]
			case r1.low <= r2.low && r1.low <= r2.hi && r1.hi >= r2.low && r1.hi >= r2.hi:
				{
					hasOverlap = true
					newRange = r1
					return
				}
			// case 4: {--[-r1-]--}
			case r1.low >= r2.low && r1.low <= r2.hi && r1.hi >= r2.low && r1.hi <= r2.hi:
				{
					hasOverlap = true
					newRange = r2
					return
				}
			// case 5: no overlap
			default:
				{
				}
			}
		}
	}

	return
}
