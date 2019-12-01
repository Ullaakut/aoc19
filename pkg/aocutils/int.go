package aocutils

import "strconv"

// Atoi gets an int from a string.
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

// SignInt returns the absolute value of an int.
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// SignInt returns the sign of an int.
func SignInt(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

// MinInt returns the smallest int between the given arguments.
func MinInt(x, y int) int {
	if y < x {
		return y
	}
	return x
}

// MaxInt returns the largest int between the given arguments.
func MaxInt(x, y int) int {
	if y > x {
		return y
	}
	return x
}
