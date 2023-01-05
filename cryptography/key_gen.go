package cryptography

import (
	"crypto/rand"
	"log"
	"math/big"
)

func findE(totient *big.Int) *big.Int {
	e := big.NewInt(int64(1<<16 + 1))
	for e.Cmp(totient) < 0 {

		if new(big.Int).GCD(nil, nil, e, totient).Cmp(big.NewInt(1)) == 0 {
			return e
		}
		e.Add(e, big.NewInt(1))
	}
	log.Fatalln("Could not find a valid e")
	return nil
}

func GenerateKeys() *RSAkey {
	// Fist we need to generate two prime numbers
	// for n to have 2048 bits, we need to generate primes with 1024 bits
	// Generate the first prime number
	p, err := rand.Prime(rand.Reader, 1024)
	// Generate the second prime number
	q, err := rand.Prime(rand.Reader, 1024)
	if err != nil {
		log.Fatalln(err)
	}

	// Calculate n
	n := new(big.Int).Mul(p, q)

	// Calculate the totient
	pMinusOne := new(big.Int).Sub(p, big.NewInt(1))
	qMinusOne := new(big.Int).Sub(q, big.NewInt(1))
	PHI := new(big.Int).Mul(pMinusOne, qMinusOne)

	totient := new(big.Int).Div(PHI, new(big.Int).GCD(nil, nil, pMinusOne, qMinusOne))

	// find a number e such that 1 < e < totient and gcd(e, totient) = 1
	e := findE(totient)

	// find a number d such that d*e = 1 mod Totient
	d := new(big.Int)
	inverse := d.ModInverse(e, totient)
	if inverse == nil {
		log.Fatalln("Could not find a valid d")
	}

	//Here we save informations to be used on the decryption process
	dp := new(big.Int).Mod(d, pMinusOne)
	dq := new(big.Int).Mod(d, qMinusOne)

	qinv := new(big.Int)
	inverse = qinv.ModInverse(q, p)

	if inverse == nil {
		log.Fatalln("Could not find a valid qinv")
	}

	return NewRSAkey(e, d, n, p, q, dp, dq, qinv)
}
