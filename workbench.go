package main

func main() {
	// rand.Seed(time.Now().Unix())
	// permutation := rand.Perm(50)
	//
	// fmt.Println(permutation)
	// sort.TimSort(permutation)
	// fmt.Println(permutation)
	//	arrSz := 10000
	// arrSz := 10
	// mid := arrSz / 2
	// index := 0
	// array := make([]int, arrSz)
	// for i := 0; i < arrSz; i++ {
	// 	if i%2 == 0 {
	// 		array[index] = i
	// 		index++
	// 	} else {
	// 		array[mid] = i
	// 		mid++
	// 	}
	// }
	// fmt.Println(array)
	//
	// array := rand.Perm(50)
	// fmt.Println(len(array))
	// //Calculate the memory used by the Slice
	// var mem runtime.MemStats
	// runtime.ReadMemStats(&mem)
	// fmt.Println("Memor used : ", mem.Alloc)
	a := []int{}

	for i := 0; i < 1000; i++ {
		a = append(a, i)
	}

}
