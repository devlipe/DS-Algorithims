package calc

// This is a simples implementation of the Euclidean algorithm
// It is used to calculate the greatest common divisor of two numbers
func GCD(a, b uint64) uint64 {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
