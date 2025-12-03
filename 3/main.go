package main

import (
	"fmt"
	io "github.com/faideww/aoc-2025/lib"
	"os"
)

func main() {
	input := io.ReadInputFile(os.Args[1])
	lines := io.TrimAndSplit(input)

	sum := sumConfigurations(lines, 2)
	fmt.Printf("sum(2):%d\n", sum)

	sum2 := sumConfigurations(lines, 12)
	fmt.Printf("sum(12):%d\n", sum2)
}

func sumConfigurations(banks []string, n int) int64 {
	sum := int64(0)
	for _, bank := range banks {
		sum += findLargestConfiguration(bank, n)
	}

	return sum
}

func findLargestConfiguration(bank string, n int) int64 {
	value := int64(0)
	lastIndex := -1
	for m := n; m > 0; m-- {
		d, nextIndex := findNextLargestDigit(bank, lastIndex, len(bank)-(m-1))
		value += io.PowInt64(10, m-1) * int64(d)
		lastIndex = nextIndex
	}
	return value
}

func findNextLargestDigit(bank string, offset int, limit int) (int, int) {
	largest := 0
	largestIndex := -1
	for j := offset + 1; j < limit; j++ {
		value := digit(bank[j])
		if value > largest {
			largest = value
			largestIndex = j
		}
	}

	return largest, largestIndex
}

func digit(r byte) int {
	switch r {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		panic("not a digit")
	}
}
