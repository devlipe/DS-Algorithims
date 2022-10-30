package main

import (
	"bytes"
	"fmt"
	"log"

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
	s := `RSA (Rivest–Shamir–Adleman) is a public-key cryptosystem that is widely used for secure data transmission. It is also one of the oldest. The acronym "RSA" comes from the surnames of Ron Rivest, Adi Shamir and Leonard Adleman, who publicly described the algorithm in 1977. An equivalent system was developed secretly in 1973 at GCHQ (the British signals intelligence agency) by the English mathematician Clifford Cocks. That system was declassified in 1997.[1]

In a public-key cryptosystem, the encryption key is public and distinct from the decryption key, which is kept secret (private). An RSA user creates and publishes a public key based on two large prime numbers, along with an auxiliary value. The prime numbers are kept secret. Messages can be encrypted by anyone, via the public key, but can only be decoded by someone who knows the prime numbers.[2]

The security of RSA relies on the practical difficulty of factoring the product of two large prime numbers, the "factoring problem". Breaking RSA encryption is known as the RSA problem. Whether it is as difficult as the factoring problem is an open question.[3] There are no published methods to defeat the system if a large enough key is used.

RSA is a relatively slow algorithm. Because of this, it is not commonly used to directly encrypt user data. More often, RSA is used to transmit shared keys for symmetric-key cryptography, which are then used for bulk encryption–decryption. `
	sb := []byte(s)
	key := cryptography.GenerateKeys()

	key.Print()
	fmt.Println("Original: ", s)
	encrypted := key.Encrypt(sb)

	// fmt.Println("encoded> ", encrypted)

	decrypted := key.Decrypt(encrypted)

	// fmt.Println("decoded> ", decrypted)
	fmt.Println("Decrypted: ", bytes.NewBuffer(decrypted).String())
	if bytes.Equal(sb, decrypted) {
		fmt.Println("Equal")
	} else {
		log.Fatal("Not Equal")
	}
}
