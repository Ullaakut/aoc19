package aocutils

import "strconv"

// Atof gets a float from a string.
func Atof(s string) float64 {
	i, err := strconv.ParseFloat(s, 10)
	Check(err)
	return i
}

// AbsFloat returns the absolute value of a float.
func AbsFloat(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// SignFloat returns the sign of a float.
func SignFloat(x float64) float64 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

// MinFloat returns the smallest float between the given arguments.
func MinFloat(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}

// MinFloat returns the largest float between the given arguments.
func MaxFloat(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}
