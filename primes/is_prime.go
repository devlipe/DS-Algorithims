package primes

import "math/big"

func IsPrime(n *big.Int) bool {
	if n.Cmp(big.NewInt(2)) == 0 || n.Cmp(big.NewInt(3)) == 0 {
		return true
	}
	nmodtwo := new(big.Int).Mod(n, big.NewInt(2))
	nmodthree := new(big.Int).Mod(n, big.NewInt(3))
	if n.Cmp(big.NewInt(1)) <= 0 || nmodtwo.Cmp(big.NewInt(0)) == 0 || nmodthree.Cmp(big.NewInt(0)) == 0 {
		return false
	}
	var i big.Int
	sqrtofn := new(big.Int).Sqrt(n)

	for i = *big.NewInt(5); i.Cmp(sqrtofn) <= 0; i.Add(&i, big.NewInt(6)) {

		nmodi := new(big.Int).Mod(n, &i)
		nmodiplus2 := new(big.Int).Mod(n, new(big.Int).Add(&i, big.NewInt(2)))

		if nmodi.Cmp(big.NewInt(0)) == 0 || nmodiplus2.Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}

	return true
}
