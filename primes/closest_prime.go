package primes

import (
	"math/big"
)

func NextPrime(n *big.Int) *big.Int {
	i := new(big.Int)
	for i.Add(n, big.NewInt(1)); ; i.Add(i, big.NewInt(1)) {

		// We coded the miler rabbin test to return true if the number is NextPrime
		// but go provided the test inside the package math/big
		// so lets use it

		if !i.ProbablyPrime(100) {
			continue
		}
		// This will take forever to run
		// So we will not use it, but its implemented

		if IsPrime(i) {
			return i
		}
	}
}

func PreviousPrime(n *big.Int) *big.Int {
	i := new(big.Int)
	for i.Sub(n, big.NewInt(1)); ; i.Sub(i, big.NewInt(1)) {

		if !i.ProbablyPrime(100) {
			continue
		}
		if IsPrime(i) {
			return i
		}
	}
}
