package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	io "github.com/faideww/aoc-2025/lib"
)

func main() {
	input := io.ReadInputFile(os.Args[1])
	ranges := io.TrimAndSplitBy(input, ",")

	count := sumInvalidRanges(ranges)

	fmt.Printf("sum:%d\n", count)
}

func sumInvalidRanges(ranges []string) int {
	count := 0
	for _, r := range ranges {
		count += sumInvalidIds(r)
	}

	return count
}

func sumInvalidIds(r string) int {
	count := 0
	components := strings.Split(r, "-")
	low := components[0]
	hi := components[1]

	lowInt, err := strconv.Atoi(low)
	if err != nil {
		panic(err)
	}

	hiInt, err := strconv.Atoi(hi)
	if err != nil {
		panic(err)
	}

	for current := lowInt; current <= hiInt; current++ {
		currentStr := fmt.Sprintf("%d", current)
		strLen := len(currentStr)

		for splitLen := strLen / 2; splitLen > 0; splitLen-- {
			if strLen%splitLen == 0 {
				lastSplit := currentStr[:splitLen]
				allEqual := true
				for i := 0; i < strLen/splitLen; i++ {
					currentSplit := currentStr[splitLen*i : splitLen*(i+1)]
					if currentSplit != lastSplit {
						allEqual = false
						break
					}
				}

				if allEqual {
					count += current
					// break so we don't count the same id multiple times
					break
				}
			}
		}

	}

	return count
}
