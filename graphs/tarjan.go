package graphs

import "golang.org/x/exp/constraints"

// Tarjan algorithm for finding bridges and articulations in a graph
//in a directed graph.

// Pair struct with generich types
type Pair[T, U any] struct {
	First  T
	Second U
}

//Create a min function that return the minimum of two integers
func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Tarjan(graph [][]int, articulations []bool, bridges []Pair[int, int]) {
	sz := len(graph)

	articulations = make([]bool, sz)        // articulations must be initialized to false
	bridges = make([]Pair[int, int], 0, sz) // the size of the array is 0 because len(bridges) is the number of bridges on the graph

	father := make([]int, sz) // father must be initialized to -1
	for i := range father {
		father[i] = -1
	}

	nivelBaixo := make([]int, sz)
	nivel := make([]int, sz)
	contadorNivel := 0

	for i := 0; i < sz; i++ {
		if nivel[i] == 0 {
			filhosDaRaiz := 0
			TarjanRecursivo(i, graph, father, nivel, nivelBaixo, articulations, bridges, &contadorNivel, &filhosDaRaiz, i)
			articulations[i] = filhosDaRaiz > 1
		}
	}
}

func TarjanRecursivo(i int, graph [][]int, father []int, nivel []int, nivelBaixo []int, articulations []bool, bridges []Pair[int, int], contadorNivel *int, filhosDaRaiz *int, nivelRaiz int) {
	*contadorNivel++
	nivel[i] = *contadorNivel
	nivelBaixo[i] = *contadorNivel

	for _, v := range graph[i] {
		if nivel[v] == 0 {
			father[v] = i

			if i == nivelRaiz {
				*filhosDaRaiz++
			}

			TarjanRecursivo(v, graph, father, nivel, nivelBaixo, articulations, bridges, contadorNivel, filhosDaRaiz, nivelRaiz)

			if nivelBaixo[v] >= nivel[i] {
				articulations[i] = true
				if nivelBaixo[v] > nivel[i] {

					if v < i {
						bridges = append(bridges, Pair[int, int]{v, i})
					} else {
						bridges = append(bridges, Pair[int, int]{i, v})
					}
					nivelBaixo[i] = min(nivelBaixo[v], nivelBaixo[i])
				}
			}
			nivelBaixo[i] = min(nivelBaixo[v], nivelBaixo[i])
		} else if v != father[i] {
			nivelBaixo[i] = min(nivelBaixo[v], nivelBaixo[i])
		}
	}
}
