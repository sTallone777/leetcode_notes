/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	c := 0
	for i, v := range nums {
		if v != val {
			if c > 0 {
				nums[i-c] = v
			}
		} else {
			c++
		}
	}
	return len(nums) - c
}

// @lc code=end

