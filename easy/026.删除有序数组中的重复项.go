func removeDuplicates(nums []int) int {
    // m:=make(map[int]int)
    // for i:=0; i<len(nums); {
    //     if _, ok := m[nums[i]]; ok{
    //         if i < len(nums) -1 {
    //             nums = append(nums[:i], nums[i+1:]...)  
    //         }else{
    //             nums = nums[:i]
    //         }
    //     }else{
    //         m[nums[i]] = i
    //         i++
    //     }
    // }
    // return len(nums)
    if len(nums) < 2 {
        return len(nums)
    }
    j := 0
    for i:=1; i<len(nums);i++{
        if nums[j] != nums[i]{
            j++
            nums[j] = nums[i]
        }
    }
    return len(nums[:j+1])
}