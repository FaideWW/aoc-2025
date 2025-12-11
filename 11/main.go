package main

import (
	"fmt"
	"os"
	"strings"

	io "github.com/faideww/aoc-2025/lib"
)

type Graph struct {
	nodes    map[string][]string
	inverted map[string][]string
}

func main() {
	input := io.ReadInputFileNoTrim(os.Args[1])
	lines := io.TrimAndSplit(input)

	graph := parseGraph(lines)

	paths := countAllPaths(graph)
	fmt.Printf("paths:%d\n", paths)

	waypointedPaths := countPathsWaypointed(graph)
	fmt.Printf("waypointed:%d\n", waypointedPaths)
}

func parseGraph(lines []string) Graph {
	graph := Graph{}
	graph.nodes = make(map[string][]string)
	graph.inverted = make(map[string][]string)

	for _, line := range lines {
		id := line[0:3]
		outs := strings.Fields(line[4:])

		graph.nodes[id] = make([]string, len(outs))
		copy(graph.nodes[id], outs)

		for _, out := range outs {
			if _, ok := graph.inverted[out]; !ok {
				graph.inverted[out] = []string{}
			}
			graph.inverted[out] = append(graph.inverted[out], id)
		}
	}

	return graph
}

func countAllPaths(g Graph) int {
	pathCache := make(map[string]int)
	return countPaths(g, "you", "out", pathCache)
}

func countPaths(g Graph, originId string, nodeId string, pathCache map[string]int) int {
	if nodeId == originId {
		pathCache[nodeId] = 1
	}

	if _, ok := pathCache[nodeId]; !ok {
		sum := 0
		for _, n := range g.inverted[nodeId] {
			sum += countPaths(g, originId, n, pathCache)
		}

		pathCache[nodeId] = sum
	}

	return pathCache[nodeId]
}

func countPathsWaypointed(g Graph) int {
	count1 := 1
	pathCache := make(map[string]int)
	count1 *= countPaths(g, "svr", "fft", pathCache)
	pathCache = make(map[string]int)
	count1 *= countPaths(g, "fft", "dac", pathCache)
	pathCache = make(map[string]int)
	count1 *= countPaths(g, "dac", "out", pathCache)

	count2 := 1
	pathCache = make(map[string]int)
	count2 *= countPaths(g, "svr", "dac", pathCache)
	pathCache = make(map[string]int)
	count2 *= countPaths(g, "dac", "fft", pathCache)
	pathCache = make(map[string]int)
	count2 *= countPaths(g, "fft", "out", pathCache)

	return count1 + count2
}
