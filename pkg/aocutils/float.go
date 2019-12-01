package aocutils

func AbsFloat(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func SignFloat(x float64) float64 {
	if x > 0 {
		return 1
	}
	if x < 0 {
		return -1
	}
	return 0
}

func MinFloat(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}

func MaxFloat(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}
