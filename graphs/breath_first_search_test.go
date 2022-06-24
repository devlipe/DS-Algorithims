package graphs_test

import (
	"testing"

	"github.com/devlipe/data_structures/graphs"
	st "github.com/devlipe/data_structures/structures"
)

func TestBfs(t *testing.T) {
	graph := st.NewDigraph[int]()

	graph.AddEdge(1, 3, 0)
	graph.AddEdge(1, 2, 0)
	graph.AddEdge(3, 8, 0)
	graph.AddEdge(2, 12, 0)
	graph.AddEdge(12, 13, 0)
	graph.AddEdge(13, 14, 0)

	var result int
	walks := 0
	visited := make(map[int]bool)
	order := make(map[int]int)
	graphs.Bfs(graph, &visited, 1, &order, func(v int, stop *bool) {
		walks++
		if i := v; i > 10 && i%2 != 0 {
			result = v
			*stop = true
		}
	})

	if result != 13 {
		t.Error("bad result")
	}

	if walks != 6 {
		t.Error("should visit 6 vertices")
	}
}
