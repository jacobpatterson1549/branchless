package branchless

import (
	"testing"
	"unsafe"
)

func TestIntSize(t *testing.T) {
	var i int
	var got = unsafe.Sizeof(i)
	const want_bits = 64
	const bits_per_byte = 8
	const want uintptr = want_bits / bits_per_byte
	if want != got {
		t.Errorf("wanted int size to be %v, got %v", want, got)
	}
}

func FuzzAll(f *testing.F) {
	f.Add(0, 0, 0)
	f.Fuzz(func(t *testing.T, x int, y int, z int) {
		var want, got int

		// Min
		want = y
		if x < y {
			want = x
		}
		got = Min(x, y)
		if want != got {
			xyError(t, "Min", x, y, want, got)
		}

		// Max
		want = y
		if x > y {
			want = x
		}
		got = Max(x, y)
		if want != got {
			xyError(t, "Max", x, y, want, got)
		}

		// MaxZeroAnd
		want = 0
		if x > 0 {
			want = x
		}
		got = MaxZeroAnd(x)
		if want != got {
			xError(t, "MaxZeroAnd", x, want, got)
		}

		// LessThan
		want = 0
		if x < y {
			want = 1
		}
		got = LessThan(x, y)
		if want != got {
			xyError(t, "LessThan", x, y, want, got)
		}

		// GreaterThan
		want = 0
		if x > y {
			want = 1
		}
		got = GreaterThan(x, y)
		if want != got {
			xyError(t, "GreaterThan", x, y, want, got)
		}

		// Equal
		want = 0
		if x == y {
			want = 1
		}
		got = Equal(x, y)
		if want != got {
			xyError(t, "Equal", x, y, want, got)
		}

		// NotEqual
		want = 0
		if x != y {
			want = 1
		}
		got = NotEqual(x, y)
		if want != got {
			xyError(t, "NotEqual", x, y, want, got)
		}

		// LessThanEqualTo
		want = 0
		if x <= y {
			want = 1
		}
		got = LessThanEqualTo(x, y)
		if want != got {
			xyError(t, "LessThanEqualTo", x, y, want, got)
		}

		// GreaterThanEqualTo
		want = 0
		if x >= y {
			want = 1
		}
		got = GreaterThanEqualTo(x, y)
		if want != got {
			xyError(t, "GreaterThanEqualTo", x, y, want, got)
		}

		// Abs
		want = x
		if want < 0 {
			want *= -1
		}
		got = Abs(x)
		if want != got {
			xError(t, "Abs", x, want, got)
		}

		// Diff
		want = x - y
		if want < 0 {
			want *= -1
		}
		got = Diff(x, y)
		if want != got {
			xyError(t, "Diff", x, y, want, got)
		}

		// Sign
		want = 0
		if x > 0 {
			want = 1
		}
		if x < 0 {
			want = -1
		}
		got = Sign(x)
		if want != got {
			xError(t, "Sign", x, want, got)
		}

		// IsPositive
		want = 0
		if x > 0 {
			want = 1
		}
		got = IsPositive(x)
		if want != got {
			xError(t, "IsPositive", x, want, got)
		}

		// IsNegative
		want = 0
		if x < 0 {
			want = 1
		}
		got = IsNegative(x)
		if want != got {
			xError(t, "IsNegative", x, want, got)
		}

		// IsZero
		want = 0
		if x == 0 {
			want = 1
		}
		got = IsZero(x)
		if want != got {
			xError(t, "IsZero", x, want, got)
		}

		// IsNotZero
		want = 0
		if x != 0 {
			want = 1
		}
		got = IsNotZero(x)
		if want != got {
			xError(t, "IsNotZero", x, want, got)
		}

		// Clamp
		var min, max = y, z
		if min > max {
			min, max = max, min
		}
		want = x
		if x < min {
			want = min
		}
		if x > max {
			want = max
		}
		got = Clamp(x, min, max)
		if want != got {
			xyzError(t, "Clamp", x, min, max, want, got)
		}

		// IsPowerOfTwo
		want = 0
		if x > 0 && x&(x-1) == 0 {
			want = 1
		}
		got = IsPowerOfTwo(x)
		if want != got {
			xError(t, "IsPowerOfTwo", x, want, got)
		}

		// IsEven
		want = 0
		if x%2 == 0 {
			want = 1
		}
		got = IsEven(x)
		if want != got {
			xError(t, "IsEven", x, want, got)
		}

		// IsOdd
		want = 0
		if x%2 != 0 {
			want = 1
		}
		got = IsOdd(x)
		if want != got {
			xError(t, "IsOdd", x, want, got)
		}

		// Negate
		want = -x
		got = Negate(x)
		if want != got {
			xError(t, "Negate", x, want, got)
		}
	})
}

func xError(t *testing.T, name string, x, want, got int) {
	t.Helper()
	t.Errorf("%v(%v) - wanted %v, got %v", name, x, want, got)
}

func xyError(t *testing.T, name string, x, y, want, got int) {
	t.Helper()
	t.Errorf("%v(%v, %v) - wanted %v, got %v", name, x, y, want, got)
}

func xyzError(t *testing.T, name string, x, y, z, want, got int) {
	t.Helper()
	t.Errorf("%v(%v, %v, %v) - wanted %v, got %v", name, x, y, z, want, got)
}
