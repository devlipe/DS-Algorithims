package structures

import "container/heap"

type Item struct {
	value    interface{} // Can be any type
	priority int         // The priority of the item in the queue
	index    int         // The index of the element in the heap
}

type PriorityQueue []*Item // Hold the itens and implements the heap interface

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want to have the highest priority element first, so we use greater than here.

	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // memory leak
	item.index = -1 // not in the queue anymore
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, value any, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
