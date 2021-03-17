/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	ret := -1
	l := len(haystack)
	for i, v := range haystack {
		if v == rune(needle[0]) {
			ret = i
			for j, n := range needle {

				if i+j >= l || rune(haystack[i+j]) != n {
					ret = -1
					break
				}
			}
		}

		if ret != -1 {
			break
		}
	}

	return ret
}

// @lc code=end

