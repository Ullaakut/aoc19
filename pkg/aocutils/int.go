package aocutils

import "strconv"

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func SignInt(x int) int {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func MinInt(x, y int) int {
	if y < x {
		return y
	}
	return x
}

func MaxInt(x, y int) int {
	if y > x {
		return y
	}
	return x
}
