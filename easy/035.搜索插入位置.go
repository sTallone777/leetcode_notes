package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ret := len(nums)
	for left <= right {
		tmp := left + (right-left)>>1
		if target <= nums[tmp] {
            ret = tmp
			right = tmp - 1
		} else {
			left = tmp + 1
		} 
	}
	return ret
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0
	fmt.Printf("search insert: %d\r\n", searchInsert(nums, target))
}
