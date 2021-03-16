func isIsomorphic(s string, t string) bool {
	for k, v := range s {
		if strings.IndexRune(s, v) != strings.IndexRune(t, rune(t[k])) {
			return false
		}
	}
	return true
}