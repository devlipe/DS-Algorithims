package cryptography

type RSAkey struct {
	e, d, n, p, q, dp, dq, qinv uint64
}

// Create a new key
func NewRSAkey(e, d, n, p, q, dp, dq, qinv uint64) *RSAkey {
	return &RSAkey{e, d, n, p, q, dp, dq, qinv}
}
