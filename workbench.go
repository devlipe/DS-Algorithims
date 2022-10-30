package main

import (
	"bytes"
	"fmt"

	"github.com/devlipe/data_structures/cryptography"
)

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
	// a := []int{}
	//
	// for i := 0; i < 1000; i++ {
	// 	a = append(a, i)
	// }

	// var a uint64 = 5000000
	// var b uint64 = 4562415854
	// fmt.Println(primes.PreviousPrime(b))
	// for i := 0; a != 2; i++ {
	// 	a = primes.NextPrime(a)
	// 	fmt.Println(a)
	// }

	// fmt.Println(calc.ModInverse(3, 26))
	//
	s := "Esse é um texto de teste para a criptografia RSA"
	sb := []byte(s)
	key := cryptography.GenerateKeys()

	key.Print()
	fmt.Println("Original: ", sb)
	encrypted := key.Encrypt(sb)

	fmt.Println("encoded> ", encrypted)

	decrypted := key.Decrypt(encrypted)

	fmt.Println("decoded> ", decrypted)
	fmt.Println("Decrypted: ", bytes.NewBuffer(decrypted).String())
}
