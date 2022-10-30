package calc

// Extended Euclid algorithm
func ExtendedEuclid(a, b uint64) (uint64, uint64, uint64) {
	x, y, u, v := uint64(0), uint64(1), uint64(1), uint64(0)

	for a != 0 {
		q, r := b/a, b%a
		m, n := x-u*q, y-v*q
		b, a, x, y, u, v = a, r, u, v, m, n
	}
	return b, x, y // b = gcd(a, b)
}
