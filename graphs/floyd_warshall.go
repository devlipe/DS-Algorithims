package graphs

import st "github.com/devlipe/data_structures/structures"

// FloydWarshall algorithm for finding shortest paths in a graph, it finds the shortest path between any two vertices
// Time complexity of the algorithm is O(V^3)
// Space complexity of the algorithm is O(V^2)
func FloydWarshall[T st.Vertex](graph *st.Graph[T]) map[T]map[T]float64 {
	mDist := make(map[T]map[T]float64)
	graph.EachVertex(func(v1 T, _ func()) {
		mDist[v1] = make(map[T]float64)

		graph.EachHalfEdge(v1, func(edge st.HalfEdge[T], _ func()) {
			mDist[v1][edge.End] = edge.Cost
		})
	})

	graph.EachVertex(func(k T, _ func()) {
		graph.EachVertex(func(i T, _ func()) {
			graph.EachVertex(func(j T, _ func()) {
				_, ok2 := mDist[i][k]
				_, ok3 := mDist[k][j]

				if !ok2 || !ok3 {
					return // skip if one of the vertices is not in the graph
				}

				if _, ok := mDist[i][j]; ok {

					// If there is a shortest path between i and k and k and j, then the path between i and j is updated
					mDist[i][j] = min(mDist[i][j], mDist[i][k]+mDist[k][j])
					// We can also use min(..., max(...)) to minimize the maximum edge weight
					// Or use max(..., min(...)) to maximize the minimum edge weight
				} else {
					mDist[i][j] = mDist[i][k] + mDist[k][j]
				}

			})
		})
	})
	return mDist
}
