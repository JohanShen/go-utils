package utils

func Int64Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func Int64Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func IntMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func IntMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func IfElseInt(condition bool, valt int, valf int) int {
	if condition {
		return valt
	}
	return valf
}
