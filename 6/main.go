package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	io "github.com/faideww/aoc-2025/lib"
)

type Equation struct {
	values   []int
	operator string
}

func main() {
	input := io.ReadInputFileNoTrim(os.Args[1])

	equations := parseEquations(input)

	sum := sumEquations(equations)
	fmt.Printf("sum:%d\n", sum)

	// Part 2

	equations2 := parseEquations2(input)
	sum2 := sumEquations(equations2)
	fmt.Printf("sum2:%d\n", sum2)
}

func parseEquations(input string) []Equation {
	lines := io.TrimAndSplit(input)
	values := make([][]string, len(lines))
	for i, line := range lines {
		values[i] = strings.Fields(line)
	}
	eqs := make([]Equation, len(values[0]))

	for i := range eqs {
		eq := Equation{
			values: make([]int, len(lines)-1),
		}
		for r := range len(values) - 1 {
			v, err := strconv.Atoi(values[r][i])
			if err != nil {
				panic(err)
			}
			eq.values[r] = v
		}

		eq.operator = values[len(values)-1][i]

		eqs[i] = eq
	}

	return eqs
}

func sumEquations(eqs []Equation) int {
	sum := 0
	for _, eq := range eqs {
		sum += solveEquation(eq)
	}
	return sum
}

func solveEquation(eq Equation) int {
	switch eq.operator {
	case "+":
		{
			result := 0
			for _, v := range eq.values {
				result += v
			}
			return result
		}
	case "*":
		{
			result := 1
			for _, v := range eq.values {
				result *= v
			}
			return result
		}
	}
	return -1
}

func parseEquations2(input string) []Equation {
	// We need to preserve the whitespace in this part, so we
	// can't use our usual utilities
	lines := strings.Split(input, "\n")
	// Detect an empty line at the end of the file and trim it
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	opsStr := lines[len(lines)-1]
	maxLength := len(opsStr)
	ops := []byte{}

	type Width struct {
		offset int
		width  int
	}
	widths := []Width{}
	lastOperator := 0
	ops = append(ops, opsStr[0])
	for i := 1; i < maxLength; i++ {
		if i < len(opsStr) && opsStr[i] != ' ' {
			widths = append(widths, Width{
				offset: lastOperator,
				width:  i - 1 - lastOperator,
			})
			lastOperator = i
			ops = append(ops, opsStr[i])
		}
	}
	widths = append(widths, Width{
		offset: lastOperator,
		width:  maxLength - lastOperator,
	})

	eqs := make([]Equation, len(widths))
	for i, w := range widths {
		eq := Equation{}
		for x := 0; x < w.width; x++ {
			value := 0
			currentPlace := 0
			for y := len(lines) - 2; y >= 0; y-- {
				c := byte(' ')
				if x+w.offset < len(lines[y]) {
					c = lines[y][x+w.offset]
				}
				if c != ' ' {
					digit := charToDigit(c)
					value += digit * io.PowInt(10, currentPlace)
					currentPlace++
				}
			}

			eq.values = append(eq.values, value)
		}

		eq.operator = string(ops[i])
		eqs[i] = eq
	}

	return eqs
}

func charToDigit(c byte) int {
	switch c {
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
		panic("char is not a digit")
	}
}
