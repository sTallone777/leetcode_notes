func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		idxm := m - 1
		idxn := n - 1
		if idxm >= 0 && idxn >= 0 {
			if nums1[idxm] >= nums2[idxn] {
				nums1[i] = nums1[idxm]
				m--
			} else {
				nums1[i] = nums2[idxn]
				n--
			}
		} else {
			if idxm >= 0 {
				nums1[i] = nums1[idxm]
				m--
			} else {
				nums1[i] = nums2[idxn]
				n--
			}
		}
	}
}
