package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	io "github.com/faideww/aoc-2025/lib"
)

type Button struct {
	mask   int
	values []int
}

type Machine struct {
	lights    int
	lightSize int
	buttons   []Button
	joltage   []int
}

func main() {
	input := io.ReadInputFileNoTrim(os.Args[1])
	lines := io.TrimAndSplit(input)

	machines := parseMachines(lines)
	sum := 0
	for _, m := range machines {
		// printMachine(m)
		seqLen := findLightSequence(m)
		// fmt.Printf(" sequence: %d\n", seqLen)
		sum += seqLen
	}

	fmt.Printf("part1:%d\n", sum)

	sum2 := 0
	for _, m := range machines {
		// printMachine(m)
		seqLen := findJoltageSequence(m)
		// fmt.Printf(" sequence: %d\n", seqLen)
		sum2 += seqLen
	}

	fmt.Printf("part2:%d\n", sum2)

}

func parseMachines(lines []string) []Machine {
	machines := make([]Machine, len(lines))
	for i, line := range lines {
		m := Machine{}
		components := strings.Split(line, " ")
		lights := components[0]
		buttons := components[1 : len(components)-1]
		joltage := components[len(components)-1]

		m.lights = 0
		m.lightSize = len(lights) - 2
		for j := 0; j < m.lightSize; j++ {
			if lights[j+1] == '#' {
				m.lights += io.PowInt(2, j)
			}
		}

		m.buttons = make([]Button, len(buttons))
		for j := 0; j < len(buttons); j++ {
			buttonStr := buttons[j]
			wireds := strings.Split(buttonStr[1:len(buttonStr)-1], ",")
			button := Button{
				mask:   0,
				values: make([]int, len(wireds)),
			}
			for k, w := range wireds {
				buttonId, err := strconv.Atoi(w)
				if err != nil {
					panic(err)
				}
				button.mask += io.PowInt(2, buttonId)
				button.values[k] = buttonId
			}

			m.buttons[j] = button
		}

		joltages := strings.Split(joltage[1:len(joltage)-1], ",")
		m.joltage = make([]int, len(joltages))
		for j := 0; j < len(joltages); j++ {
			value, err := strconv.Atoi(joltages[j])
			if err != nil {
				panic(err)
			}
			m.joltage[j] = value
		}

		machines[i] = m
	}
	return machines
}

func findLightSequence(m Machine) int {
	type Node struct {
		lights   int
		distance int
	}

	seen := make(map[int]struct{})
	start := Node{
		lights:   m.lights,
		distance: 0,
	}
	frontier := []Node{
		start,
	}
	seen[start.lights] = struct{}{}

	for len(frontier) > 0 {
		// fmt.Printf("%v\n", frontier)
		current := frontier[0]
		frontier = frontier[1:]

		if current.lights == 0 {
			return current.distance
		}

		for _, button := range m.buttons {
			next := Node{
				lights:   current.lights ^ button.mask,
				distance: current.distance + 1,
			}

			if _, ok := seen[next.lights]; !ok {
				frontier = append(frontier, next)
				seen[next.lights] = struct{}{}
			}

		}
	}

	panic("no sequence found")
}

func findJoltageSequence(m Machine) int {
	type Node struct {
		joltage  []int
		distance int
	}

	seen := make(map[string]struct{})
	start := Node{
		joltage:  m.joltage,
		distance: 0,
	}
	frontier := []Node{
		start,
	}
	seen[joltageToString(start.joltage)] = struct{}{}

	fmt.Printf("searching from %v\n", start)
	for len(frontier) > 0 {
		current := frontier[0]
		frontier = frontier[1:]

		done := true
		for _, j := range current.joltage {
			if j > 0 {
				done = false
				break
			}
		}
		if done {
			return current.distance
		}

		for _, button := range m.buttons {
			nextJoltage := make([]int, len(current.joltage))
			copy(nextJoltage, current.joltage)

			valid := true
			for _, b := range button.values {
				nextJoltage[b]--
				if nextJoltage[b] < 0 {
					valid = false
				}
			}

			if !valid {
				continue
			}

			next := Node{
				joltage:  nextJoltage,
				distance: current.distance + 1,
			}

			if _, ok := seen[joltageToString(next.joltage)]; !ok {
				frontier = append(frontier, next)
				seen[joltageToString(next.joltage)] = struct{}{}
			}

		}
	}

	panic("no sequence found")
}

func joltageIsEqual(j1 []int, j2 []int) bool {
	if len(j1) != len(j2) {
		return false
	}
	for i := range j1 {
		if j1[i] != j2[i] {
			return false
		}
	}
	return true
}

func joltageToString(j []int) string {
	s := strings.Builder{}
	for _, v := range j {
		fmt.Fprintf(&s, "_%d", v)
	}
	return s.String()
}

func printMachine(m Machine) {
	fmt.Printf("[%d %s] (", m.lights, strconv.FormatInt(int64(m.lights), 2))
	for _, b := range m.buttons {
		fmt.Printf("(%d %s)", b.mask, strconv.FormatInt(int64(b.mask), 2))
	}
	fmt.Printf(") {%v}\n", m.joltage)
}
