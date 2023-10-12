package sorter

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

type Node struct {
	source     int
	degree     int
	neighbours []int
}

type Graph struct {
	num_v int
	nodes map[int]*Node
}

type Sorter interface {
	Sort(node_map Graph) Graph
}

func MakeGraph() Graph {
	g := Graph{nodes: make(map[int]*Node)}
	return g
}

func CreateGraph(filename string) Graph {
	g := MakeGraph()
	nodes := g.nodes

	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	max_v := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var from, to int
		if _, err := fmt.Sscanf(line, "%d %d", &from, &to); err != nil {
			log.Fatal(err)
		}
		if from > max_v {
			max_v = from
		}
		if to > max_v {
			max_v = to
		}
		if node, ok := nodes[from]; ok {
			node.neighbours = append(node.neighbours, to)
		} else {
			node = &Node{}
			node.source = from
			node.neighbours = []int{to}
			nodes[from] = node
		}
		nodes[from].degree += 1
	}

	g.num_v = max_v + 1
	return g
}

func SaveGraph(filename string, graph Graph) {

	g := graph.nodes

	var writer io.Writer

	if filename == "" {
		writer = os.Stdout
	} else {
		writer_, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		writer = writer_
	}

	values := make([]*Node, 0, len(g))
	for _, v := range g {
		values = append(values, v)
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].source < values[j].source
	})

	for _, node := range values {

		sort.IntSlice(node.neighbours).Sort()

		for _, neignbour := range node.neighbours {
			fmt.Fprintln(writer, node.source, neignbour)
		}
	}
}

func CreateSorter(sorter_type string) Sorter {
	switch sorter_type {
	case "degree":
		return DegreeSorter{}
	}
	panic("Invalid Sorter type.")
}
