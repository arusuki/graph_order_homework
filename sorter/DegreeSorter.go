package sorter

import (
	"math"
)

type DegreeSorter struct {
}

type void_t struct{}

const w = 5

var void void_t

func s(g *Graph, u, v int) int {
	v_neighbour := make(map[int]void_t)
	for _, neighbour := range (*g)[v].neighbours {
		v_neighbour[neighbour] = void
	}
	sn := 0
	ss := 0
	for _, neighbour := range (*g)[u].neighbours {
		if _, ok := v_neighbour[neighbour]; ok {
			ss += 1
		}
	}
	if _, ok := v_neighbour[u]; ok {
		sn = 2
	}
	return sn + ss
}

func (DegreeSorter) Sort(old_graph Graph) Graph {

	new_graph := make(Graph)
	V_R := make(Graph)

	for k, v := range old_graph {
		V_R[k] = v
	}

	n := len(V_R)

	v := 0

	P := make([]int, 0, len(V_R))

	delete(V_R, v)

	P = append(P, v)

	for i := 1; i < n; i += 1 {
		var vmax int
		kmax := math.MinInt
		for v := range V_R {
			j := i - w
			if j < 0 {
				j = 0
			}
			kv := 0
			for ; j < i; j += 1 {
				kv += s(&old_graph, P[j], v)
			}
			if kv > kmax {
				kmax = kv
				vmax = v
			}
		}
		P = append(P, vmax)
		delete(V_R, vmax)
	}

	for k, v := range old_graph {
		node := &Node{
			source:     P[k],
			neighbours: make([]int, 0, len(v.neighbours)),
			degree:     v.degree,
		}
		for _, neighbour := range v.neighbours {
			node.neighbours = append(node.neighbours, P[neighbour])
		}

		new_graph[node.source] = node
	}

	return new_graph
}
