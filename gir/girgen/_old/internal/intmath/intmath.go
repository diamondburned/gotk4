package intmath

func Abs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func Difference(i, j int) int {
	return Abs(i - j)
}
