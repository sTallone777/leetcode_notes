/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	if s == t {
		if len(s) < 2 {
			return true
		}
		return false
	}
	sMap := countRune(s)
	tMap := countRune(t)

	for k, v := range sMap {
		val, ok := tMap[k]
		if !ok {
			return false
		} else {
			if val != v {
				return false
			}
		}
	}

	return true
}

func countRune(s string) map[rune]int {
	runeMap := make(map[rune]int)

	for _, v := range s {
		val, ok := runeMap[v]
		if !ok {
			runeMap[v] = 1
		} else {
			runeMap[v] = val + 1
		}
	}
	return runeMap
}

// @lc code=end

