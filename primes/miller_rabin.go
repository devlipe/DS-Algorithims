package primes

import (
	"math/rand"
	"time"

	"github.com/devlipe/data_structures/calc"
)

func translate_n(n uint64) (s, d uint64) {
	m := n - 1
	// d must be an odd number
	for (m & 1) == 0 {
		s += 1  // increse the power
		m >>= 1 // Divide m by 2
	}

	d = (n - 1) / (1 << s)
	return s, d
}

func MillerRabin(n uint64, trials int) bool {

	// Threat negative numbers as not primes
	if n < 1 {
		return false
	}
	// Threat this numbers as primes
	if n == 1 || n == 2 || n == 3 {
		return true
	}

	// Even numbers are not prime
	if (n & 1) == 0 {
		return false
	}
	// We have to write the number n as 2^s * d + 1
	// We will save d and s to be used futher in the function
	s, d := translate_n(n)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < trials; i++ {
		a := uint64(rand.Int63n(int64(n)-4) + 2)
		x := calc.ModExpSqr(a, d, n) // Here we calculate a^d % n

		if x == 1 || x == (n-1) {
			// If x == 1 or x == n-1 then we can't be sure if n is prime
			continue
		}

		var r uint64 = 0
		for ; r < (s - 1); r++ {

			// Here we calculate (a^2)^r * d% n
			// But we already have a^d % n
			// So we can use the property (a^2)^r * d% n = (a^d)^2^r % n
			// For each increase in r we only have to call the function pow_mod once more
			x = calc.ModExpSqr(x, 2, n)

			if x == 1 {
				return false
			}

			if x == (n - 1) {
				break
			}
		}

		if x != (n - 1) {
			return false
		}

	}

	return true
}
