package graphs

// Tarjan algorithm for finding strongly connected components
//in a directed graph.

// Pair struct with generich types
type Pair[T, U any] struct {
	First  T
	Second U
}

func Tarjan(graph [][]int, articulations []bool, bridges []Pair[int, int]) {
	sz := len(graph)

	articulations = make([]bool, sz, 0)
	bridges = make([]Pair[int, int], 0)

	father := make([]int, sz)
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
		}
	}
}

func TarjanRecursivo(i1 int, graph [][]int, father []int, nivel []int, nivelBaixo []int, articulations []bool, bridges []Pair[int, int], int1 *int, int2 *int, i2 int) {
	panic("unimplemented")
}
