func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}
	return true
}