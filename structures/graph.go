package structures

import (
	"fmt"
	"sort"
)

// Vertex can be anything
type Vertex interface {
	comparable
}

// An edge that connects two vertices.
type Edge[T Vertex] struct {
	Start, End T
	Cost       float64
}

// An Half Edge is an edge that only store the end vertex.
// The start vertex is infered from the contextual information.
type HalfEdge[T Vertex] struct {
	End  T
	Cost float64
}

// A Graph is a set of vertices and edges stored in an adjacency list.
type Graph[T Vertex] struct {
	AdjList  map[T]*Set[HalfEdge[T]]
	directed bool
}

// Create an empty new graph
func NewGraph[T Vertex]() *Graph[T] {
	return &Graph[T]{
		AdjList:  map[T]*Set[HalfEdge[T]]{},
		directed: false,
	}
}

// Create an empty directed gruufunc NewDigraph[T Vertex]() *Graph[T] {
func NewDigraph[T Vertex]() *Graph[T] {
	graph := NewGraph[T]()
	graph.directed = true
	return graph
}

// Add a new vertex to te graph
func (g *Graph[T]) AddVertex(v T) {
	if _, exists := g.AdjList[v]; exists {
		g.AdjList[v] = NewSet[HalfEdge[T]]()
	}
}

// Add edge to the Graph. The edge is from v1 to v2 with cost c
// It also adds the vertices to the graph
func (g *Graph[T]) AddEdge(v1, v2 T, c float64) {
	// If the vertices does't exists, we try to add them
	g.AddVertex(v1)
	g.AddVertex(v2)

	g.AdjList[v1].Add(HalfEdge[T]{
		End:  v2,
		Cost: c,
	})

	if !g.directed {

		g.AdjList[v2].Add(HalfEdge[T]{
			End:  v1,
			Cost: c,
		})
	}
}

// Funt to print all edges an cost to stdout
func (g *Graph[T]) Dump() {
	g.EachEdge(func(e Edge[T], _ func()) {
		fmt.Printf("(%v, %v, %f)\n", e.Start, e.End, e.Cost)
	})
}

//Return the number of vertices in the graph
func (g *Graph[T]) NVertices() int {
	return len(g.AdjList)
}

// Return the number of edges in the graph
func (g *Graph[T]) NEdges() int {
	count := 0

	for _, v := range g.AdjList {
		count += v.Len()
	}

	// If the graph is undirected we dont count
	// a-b and b-a as different Edges
	if !g.directed {
		return count / 2
	}
	return count
}

// Tests if two given graphs are equal.
func (g *Graph[T]) Equals(other *Graph[T]) bool {
	// They are equal if they have the same number of vertices...
	if g.NVertices() != other.NVertices() {
		return false
	}

	// If they have the same amount of edges...
	if g.NEdges() != other.NEdges() {
		return false
	}

	// Now we have to test if they have the same edges
	for k, v := range g.AdjList {

		if !v.Equals(other.AdjList[k]) {
			return false
		}
	}
	return true
}

//Call a function f to each vertex
func (g *Graph[T]) EachVertex(f func(T, func())) {
	var stopped bool
	stop := func() { stopped = true }

	for k := range g.AdjList {
		f(k, stop)
		if stopped {
			break
		}
	}
}

// SortedEdges is an array of edges that can be sorted
type SortedEdges[T Vertex] []Edge[T]

// Return the length of the sorted edges
func (se SortedEdges[T]) Len() int {
	return len(se)
}

// Return if a edge is less than other
func (se SortedEdges[T]) Less(i, j int) bool {
	return se[i].Cost < se[j].Cost
}

// Function that swaps two edges
func (se SortedEdges[T]) Swap(i, j int) {
	se[i], se[j] = se[j], se[i]
}

// Return an array of sorted edges based on their cost
func (g *Graph[T]) SortedEdges() SortedEdges[T] {
	set := NewSet[Edge[T]]()

	for v := range g.AdjList {
		g.EachHalfEdge(v, func(he HalfEdge[T], _ func()) {
			set.Add(Edge[T]{
				Start: v,
				End:   he.End,
				Cost:  he.Cost,
			})
		})
	}

	edges := make(SortedEdges[T], set.Len())

	set.Each(func(e Edge[T], _ func()) {
		edges = append(edges, e)
	})

	sort.Sort(&edges) // We can do this because whe implemented Less and Swap func
	return edges
}

// EachHalfEdge call a function f to every edge of a given Vertex
func (g *Graph[T]) EachHalfEdge(v T, f func(HalfEdge[T], func())) {
	if ed, exists := g.AdjList[v]; exists {
		var stopped bool
		stop := func() { stopped = true }
		ed.Each(func(he HalfEdge[T], innerStop func()) {
			f(he, stop)
			if stopped {
				innerStop()
			}
		})
	}
}

// func that will call other function for every edge
func (g *Graph[T]) EachEdge(f func(Edge[T], func())) {
	var stopped bool
	stop := func() { stopped = true }

	for k, s := range g.AdjList {
		s.Each(func(he HalfEdge[T], innerStop func()) {
			edge := Edge[T]{Start: k, End: he.End, Cost: he.Cost}
			f(edge, stop)
			if stopped {
				innerStop()
			}
		})
		if stopped {
			break
		}
	}
}
