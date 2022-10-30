package calc

func ChineseRemainderRSA(p, q, dp, dq, qinv, c uint64) uint64 {
	m1 := ModExpSqr(c, dp, p)
	m2 := ModExpSqr(c, dq, q)
	h := (qinv * (m1 - m2)) % p
	m := (m2 + h*q) % (p * q)

	return m
}
