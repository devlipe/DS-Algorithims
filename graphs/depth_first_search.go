package graphs

import st "github.com/devlipe/data_structures/structures"

type DFSWalkerFunc[T st.Vertex] func(T, *bool)

func Dfs[T st.Vertex](g *st.Graph[T], start T, walker DFSWalkerFunc[T]) {
	visited := st.NewSet[T]()
	stop := false
	dfs_recursive(g, start, visited, &stop, walker)

}

func dfs_recursive[T st.Vertex](g *st.Graph[T], start T, visited *st.Set[T], stop *bool, walker DFSWalkerFunc[T]) {
	visited.Add(start)

	walker(start, stop)
	if *stop {
		return
	}
	g.EachHalfEdge(start, func(e st.HalfEdge[T], stoper func()) {
		if !visited.Contains(e.End) {
			dfs_recursive(g, e.End, visited, stop, walker)
			if *stop {
				stoper()
			}
		}
	})
}
