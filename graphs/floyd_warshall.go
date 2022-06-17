package graphs

// FloydWarshall algorithm for finding shortest paths in a graph, it finds the shortest path between any two vertices
// Time complexity of the algorithm is O(V^3)
// Space complexity of the algorithm is O(V^2)
func FloydWarshall(graph [][]int) [][]int {
	sz := len(graph)

	for k := 0; k < sz; k++ {
		for i := 0; i < sz; i++ {
			for j := 0; j < sz; j++ {
				// If there is a shortest path between i and k and k and j, then the path between i and j is updated
				graph[i][j] = min(graph[i][j], graph[i][k]+graph[k][j])
				// We can also use min(..., max(...)) to minimize the maximum edge weight
				// Or use max(..., min(...)) to maximize the minimum edge weight
			}
		}
	}

	return graph
}
