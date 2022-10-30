package calc

// Function to find moludar inverse of a under modulo
// m using extended Euclid Algorithm. Refer below post for details:
func ModularInverse(a, m uint64) uint64 {
	gcd, x, _ := ExtendedEuclid(a, m)
	if gcd != 1 {
		// Moduar inverse just exists if a and m are coprimes
		return 0
	}
	return (x + m) % m
}
