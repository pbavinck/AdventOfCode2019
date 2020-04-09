package math2

// IntMax returns the highest integer of the two
func IntMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// IntMin returns the lowest integer of the two
func IntMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// IntAbs returns the absolute of an integer
func IntAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// GCD greatest common divisor via Euclidean algorithm
func GCD(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple via GCD
func LCM(a, b int64, integers ...int64) int64 {
	var result, i int64
	result = a * b / GCD(a, b)

	for i = 0; i < int64(len(integers)); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

// IntPow computes a**b using binary powering algorithm
// See Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
// Source: https://groups.google.com/d/msg/golang-nuts/PnLnr4bc9Wo/z9ZGv2DYxXoJ
func IntPow(a, b int64) int64 {
	var p int64 = 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
