package sorter

import (
	"math"
)

type DegreeSorter struct {
}

type void_t struct{}

const w = 5

var void void_t

func s(graph *Graph, u, v int) int {

	g := graph.nodes

	v_neighbour := make(map[int]void_t)
	for _, neighbour := range g[v].neighbours {
		v_neighbour[neighbour] = void
	}
	sn := 0
	ss := 0
	for _, neighbour := range g[u].neighbours {
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

	new_graph := MakeGraph()
	V_R_graph := MakeGraph()

	V_R := V_R_graph.nodes

	for k, v := range old_graph.nodes {
		V_R[k] = v
	}

	n := old_graph.num_v

	v := 0

	P := make(map[int]int)

	delete(V_R, v)

	P[0] = v

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
		P[i] = vmax
		delete(V_R, vmax)
	}

	for k, v := range old_graph.nodes {
		node := &Node{
			source:     P[k],
			neighbours: make([]int, 0, len(v.neighbours)),
			degree:     v.degree,
		}
		for _, neighbour := range v.neighbours {
			node.neighbours = append(node.neighbours, P[neighbour])
		}

		new_graph.nodes[node.source] = node
	}
	new_graph.num_v = old_graph.num_v
	return new_graph
}
