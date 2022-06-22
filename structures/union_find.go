package structures

/*
	This data structure is used to solve the problem of finding the connected components of a graph.
	It's quite simple, but a very elegant solution.
*/
type UnionFind struct {
	representant []int
}

func NewUnionFind(size int) UnionFind {
	uf := UnionFind{}
	uf.representant = make([]int, size)
	// Initialize representant to itself
	for i := 0; i < size; i++ {
		uf.representant[i] = i
	}
	return uf
}

//This functioin return the representant of the element "element"
func (uf *UnionFind) FindRepresentant(element int) int {
	if uf.representant[element] == element {
		return element
	}
	//Here we are updating the path to the representant(shrinking the path)
	uf.representant[element] = uf.FindRepresentant(uf.representant[element])
	return uf.representant[element]
}

//Given element1 and element2, this function will merge the two sets
func (uf *UnionFind) MergeSets(element1, element2 int) {
	// Here we make the representant of element1 point to the representant of element2
	uf.representant[uf.FindRepresentant(element1)] = uf.FindRepresentant(element2)
}

func (uf *UnionFind) IsSameSet(element1, element2 int) bool {
	return uf.FindRepresentant(element1) == uf.FindRepresentant(element2)
}
