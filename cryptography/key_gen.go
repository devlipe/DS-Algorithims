package cryptography

import (
	"math"
	"math/rand"
	"time"

	"github.com/devlipe/data_structures/calc"
	"github.com/devlipe/data_structures/primes"
)

// Generate a uint64 from range min to max
func randUint64(min, max uint64) (uint64, uint64) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Uint64()%(max-min) + min
	b := rand.Uint64()%(max-min) + min
	return a, b
}

func findE(totient uint64) uint64 {
	var e uint64 = 65540
	for e < totient {
		if calc.GCD(e, totient) == 1 {
			return e
		}
		e++
	}
	return 0
}

func GenerateKeys() *RSAkey {
	// Fist we need to generate two prime numbers
	rand.Seed(time.Now().Unix())
	// max := uint64(math.Sqrt(math.MaxUint64))
	// threshold_p, threshold_q := randUint64(max/100, max/10)
	threshold_p, threshold_q := randUint64(1000000, 100000000)

	// Generate the first prime number
	p := primes.PreviousPrime(threshold_p)

	// Generate the second prime number
	q := primes.PreviousPrime(threshold_q)
	// Calculate the modulus
	n := p * q

	// Calculate the totient
	totient := uint64(math.Abs(float64((p-1)*(q-1)))) / calc.GCD(p-1, q-1)

	// find a number e such that 1 < e < totient and gcd(e, totient) = 1
	e := findE(totient)

	// find a number d such that d*e = 1 mod Totient
	d := calc.ModularInverse(e, totient)

	//Here we save informations to be used on the decryption process
	dp := d % (p - 1)
	dq := d % (q - 1)

	qinv := calc.ModularInverse(q, p)

	return NewRSAkey(e, d, n, p, q, dp, dq, qinv)
}
