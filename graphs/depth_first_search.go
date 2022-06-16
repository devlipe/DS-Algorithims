package graphs

func dfs(graph map[rune][]*rune, verticeDescoberto map[rune]bool, start rune) {
	verticeDescoberto[start] = true

	for _, v := range graph[start] {
		if !verticeDescoberto[*v] {
			dfs(graph, verticeDescoberto, *v)
		}
	}
}
