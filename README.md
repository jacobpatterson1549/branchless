# Branchless

Branchless is a Go package that provides branchless operations on integers.

## Functions

### Min(x, y int) int
Returns the minimum of x and y.

### Max(x, y int) int
Returns the maximum of x and y.

### MaxZeroAnd(x int) int
Returns 0 if x is less than 0, otherwise x.

### LessThan(x, y int) int
Returns 1 if x < y, otherwise 0.

### GreaterThan(x, y int) int
Returns 1 if x > y, otherwise 0.

### Equal(x, y int) int
Returns 1 if x == y, otherwise 0.

### NotEqual(x, y int) int
Returns 1 if x != y, otherwise 0.

### LessThanEqualTo(x, y int) int
Returns 1 if x <= y, otherwise 0.

### GreaterThanEqualTo(x, y int) int
Returns 1 if x >= y, otherwise 0.

### Abs(x int) int
Returns the absolute value of x.

### Diff(x, y int) int
Returns the absolute difference between x and y.

### Sign(x int) int
Returns -1 if x < 0, 0 if x == 0, and 1 if x > 0.

### IsPositive(x int) int
Returns 1 if x > 0, otherwise 0.

### IsNegative(x int) int
Returns 1 if x < 0, otherwise 0.

### IsZero(x int) int
Returns 1 if x == 0, otherwise 0.

### IsNotZero(x int) int
Returns 1 if x != 0, otherwise 0.

### Clamp(x, min, max int) int
Clamps the value x between the minimum (min) and maximum (max) values.

### IsPowerOfTwo(x int) int
Returns 1 if x is a power of 2 (or is zero), otherwise 0.

### IsEven(x int) int
Returns 1 if x is even, otherwise 0.

### IsOdd(x int) int
Returns 1 if x is odd, otherwise 0.

### Negate(x int) int
Returns the negation of x.

## Assembly Inspection

Branchless code contains no "jump" instructions.
Jump instructions usually start with a J.
The only return for each function should be the last instruction ("RET").

1. Create branchless.o
  `go tool compile -S -o branchless.o branchless.go`
1. Print assembly of branchless.o
  `go tool objdump branchless.o`
