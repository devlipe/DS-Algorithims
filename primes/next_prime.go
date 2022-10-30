package primes

func NextPrime(n uint64) uint64 {
	var i uint64
	for i = n + 1; ; i++ {

		if !MillerRabin(i, 100) {
			continue
		}
		if IsPrime(i) {
			return i
		}
	}
}

func PreviousPrime(n uint64) uint64 {
	var i uint64
	for i = n - 1; ; i-- {

		if !MillerRabin(i, 100) {
			continue
		}
		if IsPrime(i) {
			return i
		}
	}
}
