package graphs

// Function that inplements breath first search
func Bfs(graph [][]int, descobriu []bool, source int, ordem []int) {
	queue := make([]int, 0)

	queue = append(queue, source)
	descobriu[source] = true
	ordem[source] = 0

	for len(queue) > 0 {
		proximo := queue[0]                // pega o primeiro elemento da fila
		queue = queue[1:]                  // remove o primeiro elemento da fila
		for _, v := range graph[proximo] { // percorre os adjacentes do vértice atual
			if !descobriu[v] { // se o vértice não foi descobrerto ainda
				descobriu[v] = true           // marca como descobrido
				ordem[v] = ordem[proximo] + 1 // atualiza o valor do vértice
				queue = append(queue, v)      // adiciona o vértice na fila
			}
		}
	}
}
