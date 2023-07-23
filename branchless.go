// Package branchless provides branchless operations on integers.
package branchless

import "unsafe"

const (
	zero          = 0
	int_bytes     = unsafe.Sizeof(zero)
	bits_per_byte = 8
	int_bits      = int_bytes * bits_per_byte
	max           = int(int_bits) - 1
)

// Min returns the minimum of x and y.
func Min(x, y int) int {
	return y ^ ((x ^ y) & ((x - y) >> max))
}

// Max returns the maximum of x and y.
func Max(x, y int) int {
	return x ^ ((x ^ y) & ((x - y) >> max))
}

// MaxZeroAnd returns 0 if x is less than 0, otherwise x.
func MaxZeroAnd(x int) int {
	return x & ^(x >> max)
}

// LessThan returns 1 if x < y, otherwise 0.
func LessThan(x, y int) int {
	return ((x - y) >> max) & 1
}

// GreaterThan returns 1 if x > y, otherwise 0.
func GreaterThan(x, y int) int {
	return ((y - x) >> max) & 1
}

// Equal returns 1 if x == y, otherwise 0.
func Equal(x, y int) int {
	return NotEqual(x, y) ^ 1
}

// NotEqual returns 1 if x != y, otherwise 0.
func NotEqual(x, y int) int {
	return (((x ^ y) | -(x ^ y)) >> max & 1)
}

// LessThanEqualTo returns 1 if x <= y, otherwise 0.
func LessThanEqualTo(x, y int) int {
	return GreaterThan(x, y) ^ 1
}

// GreaterThanEqualTo returns 1 if x >= y, otherwise 0.
func GreaterThanEqualTo(x, y int) int {
	return LessThan(x, y) ^ 1
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	return (x + (x >> max)) ^ (x >> max)
}

// Diff returns the absolute difference between x and y.
func Diff(x, y int) int {
	return ((x - y) ^ ((x - y) >> max)) - ((x - y) >> max)
}

// Sign returns -1 if x < 0, 0 if x == 0, and 1 if x > 0.
func Sign(x int) int {
	return IsPositive(x) - IsNegative(x)
}

// IsPositive returns 1 if x > 0, otherwise 0.
func IsPositive(x int) int {
	return IsNegative(-x)
}

// IsNegative returns 1 if x < 0, otherwise 0.
func IsNegative(x int) int {
	return (x >> max) & 1
}

// IsZero returns 1 if x == 0, otherwise 0.
func IsZero(x int) int {
	return IsNotZero(x) ^ 1
}

// IsNotZero returns 1 if x != 0, otherwise 0.
func IsNotZero(x int) int {
	return (((x | -x) >> max) & 1)
}

// Clamp clamps the value x between the minimum (min) and maximum (max) values.
func Clamp(x, min, max int) int {
	return Min(Max(x, min), max)
}

// IsPowerOfTwo returns 1 if x is a power of 2 (or is zero), otherwise 0.
func IsPowerOfTwo(x int) int {
	return IsPositive(x) & IsZero(x&(x-1))
}

// IsEven returns 1 if x is even, otherwise 0.
func IsEven(x int) int {
	return IsOdd(x) ^ 1
}

// IsOdd returns 1 if x is odd, otherwise 0.
func IsOdd(x int) int {
	return x & 1
}

// Negate returns the negation of x.
func Negate(x int) int {
	return -x
}
