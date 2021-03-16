func isPowerOfTwo(n int) bool {
	c := 0
	for i := 0; i < 64; i++ {
		c += (n >> i) & 0x01
	}
	if c == 1 {
		return true
	}
	return false
}