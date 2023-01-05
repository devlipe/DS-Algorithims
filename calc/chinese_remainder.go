package calc

import (
	"math/big"
)

func ChineseRemainderRSA(p, q, dp, dq, qinv, c *big.Int) uint64 {
	m1 := new(big.Int).Exp(c, dp, p)
	m2 := new(big.Int).Exp(c, dq, q)
	h := new(big.Int).Mul(new(big.Int).Sub(m1, m2), qinv)
	h.Mod(h, p)
	h.Mul(h, q)

	pq := new(big.Int).Mul(p, q)
	h.Add(h, m2)
	h.Mod(h, pq)

	return h.Uint64()
}
