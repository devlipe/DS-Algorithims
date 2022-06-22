package graphs

import (
	"sort"

	st "github.com/devlipe/data_structures/structures"
	"golang.org/x/exp/constraints"
)

/*
	Here is a simple implementation of an edge object. We will use it in the Kruskal's algorithm.
	The edge object has a from and to vertex, and a weight.
*/
type Edge[T constraints.Ordered] struct {
	From, To int
	Weight   T
}

func NewEdge[T constraints.Ordered](from int, to int, weight T) Edge[T] {
	return Edge[T]{from, to, weight}
}

func (e *Edge[T]) IsLessThan(other *Edge[T]) bool {
	return e.Weight < (other.Weight)
}

//Kruskal's algorithm is a greedy algorithm that finds the minimum spanning tree of a graph.
//It works by selecting the edge with the minimum weight that connects two trees, and adding it to the tree.
// Its complexity is O(ElogE) where E is the number of edges in the graph.
func Kruskal[T constraints.Ordered](arestas []Edge[T], numVertices int) (int, []bool) {

	usedEdges := make([]bool, len(arestas))
	sort.Slice(arestas, func(i, j int) bool {
		return arestas[i].IsLessThan(&arestas[j])
	})

	set := st.NewUnionFind(numVertices)
	sz := len(arestas)
	treeSz := numVertices - 1
	edgeCounter := 0

	for i := 0; i < sz; i++ {
		if !set.IsSameSet(arestas[i].From, arestas[i].To) {
			usedEdges[i] = true
			edgeCounter++

			if edgeCounter == treeSz {
				return edgeCounter, usedEdges
			}

			set.MergeSets(arestas[i].From, arestas[i].To)
		}
	}
	return edgeCounter, usedEdges
}
