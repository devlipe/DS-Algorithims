package graphs

import st "github.com/devlipe/data_structures/structures"

type BFSWalkerFunc[T st.Vertex] func(T, *bool)

// Function that inplements breath first search
func Bfs[T st.Vertex](graph *st.Graph[T], descobriu *map[T]bool, source T, ordem *map[T]int, walker BFSWalkerFunc[T]) {
	queue := make([]T, 0)

	queue = append(queue, source)
	(*descobriu)[source] = true
	(*ordem)[source] = 0

	for v := queue[0]; len(queue) > 0; v = queue[0] {
		queue = queue[1:]

		stop := false
		walker(v, &stop)

		if stop {
			return
		}

		(*descobriu)[v] = true

		graph.EachHalfEdge(v, func(he st.HalfEdge[T], _ func()) {
			if !(*descobriu)[he.End] {
				queue = append(queue, he.End)
			}
		})
	}
}
