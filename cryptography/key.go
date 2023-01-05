package cryptography

import "math/big"

type RSAkey struct {
	e, d, n, p, q, dp, dq, qinv *big.Int
}

// Create a new key
func NewRSAkey(e, d, n, p, q, dp, dq, qinv *big.Int) *RSAkey {
	return &RSAkey{e, d, n, p, q, dp, dq, qinv}
}
