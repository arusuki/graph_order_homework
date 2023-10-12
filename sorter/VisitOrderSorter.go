package sorter

import (
	queue "github.com/Workiva/go-datastructures/queue"
)

type VistOrderSorter struct {
}

func bfs(g Graph) []int {
	n := len(g)
	P := make([]int, n)
	q := queue.New(int64(n))
	q.Put(0)
	new_id := 0

	visited := make(map[int]bool)

	visited[0] = true

	for !q.Empty() {
		v, _ := q.Get(1)
		i := v[0].(int)
		for _, neighbour := range g[i].neighbours {
			if _, ok := visited[neighbour]; ok {
				continue
			}
			visited[neighbour] = true
			q.Put(neighbour)
		}
		P[i] = new_id
		new_id += 1
	}
	return P
}

func (VistOrderSorter) Sort(old_graph Graph) Graph {
	new_graph := make(Graph)

	P := bfs(old_graph)

	for k, v := range old_graph {
		if v.degree == 0 {
			continue
		}
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

	n := len(old_graph)

	for i := 0; i < n; i += 1 {
		if _, ok := new_graph[i]; !ok {
			new_graph[i] = &Node{source: i}
		}
	}

	return new_graph
}
