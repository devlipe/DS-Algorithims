package graphs

import "math"

// This implementation is using a matrix of adjancency, it can be rewritten using a list of adjancency

// Dijkstra algorithm for finding shortest paths in a graph, it finds the shortest path between any an give vertex and all other vertices
// It`s complexity is o(ElogV)
// It`s space complexity is o(V)
func Dijkstra(graph [][]int, start int) []int {
	sz := len(graph)            // number of vertices in the graph
	dist := make([]int, sz)     // distance from start vertex to each vertex
	visided := make([]bool, sz) // visited vertices, the ones that are already with the minimum distance from start vertex

	for i := 0; i < sz; i++ {
		dist[i] = math.MaxInt // distance from start vertex to each vertex is infinite, since we don`t know the distance from start vertex to any other vertex
		visided[i] = false    // all vertices are not visited
	}

	dist[start] = 0 // distance from start vertex to start vertex is 0

	// Now we`ll find the shortest path from start vertex to all other vertices
	for nvert := 0; nvert < sz-1; nvert++ {
		// Pick the vertices with the minimum distance on dist and that are not visited
		// in the first it, it is equal to start vertex
		mVert := minDistance(dist, visided)

		visided[mVert] = true // mark the vertex as visited

		// Update the distance of all the vertices connected to the picked vertex
		for v := 0; v < sz; v++ { // in case of a list of adjancency here we can loop throught the neighbors of the picked vertex directly
			/* We can only update if:
				 * 1. The vertex is not visited
			     * there is an edge from mVert to v
				 * 2. The distance from start vertex to the picked vertex is not infinite
				 * 3. The distance from start vertex to the picked vertex + the distance from the picked vertex to the connected vertex
					is less than the distance from start vertex to the connected vertex
			*/
			if !visided[v] && graph[mVert][v] > 0 && dist[mVert] != math.MaxInt && dist[mVert]+graph[mVert][v] < dist[v] {
				dist[v] = dist[mVert] + graph[mVert][v]
			}
		}
	}
	return dist
}

func minDistance(dist []int, visided []bool) int {
	min := math.MaxInt // minimum distance is infinite
	minIndex := 0      // index of minimum minDistanc
	for i := 0; i < len(dist); i++ {
		if !visided[i] && dist[i] < min {
			min = dist[i]
			minIndex = i
		}
	}
	return minIndex
}
