package src_calc

func Add(a, b int) int {
	return a + b
}

func Max(a, b int) (ret int) {
	ret = a
	if b > a {
		ret = b
	}
	return
}

func Min(a, b int) (ret int) {
	ret = a
	if b < a {
		ret = b
	}
	return
}

