package calc

// This function returns a^d % n
// We know that a^d = (a^2)^(d/2) if d is even.
// So we can use that in our favor to make the algorithm faster
// In this case we calculate a^2 % n and then we calculate (a^2)^2 % n
// It uses properties of modular arithmetic
func ModExpSqr(base, exponent, modulus uint64) uint64 {
	if modulus == 1 {
		return 0
	}
	if exponent == 0 {
		return 1
	}

	result := ModExpSqr(base, exponent/2, modulus)
	result = (result * result) % modulus
	if exponent&1 != 0 {
		return ((base % modulus) * result) % modulus
	}
	return result % modulus
}

func ModExp(base, exponent, modulus uint64) uint64 {
	if modulus == 1 {
		return 0
	}
	base = base % modulus
	result := uint64(1)
	for i := uint64(0); i < exponent; i++ {
		result = (result * base) % modulus
	}
	return result
}
