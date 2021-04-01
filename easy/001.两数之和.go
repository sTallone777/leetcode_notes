func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    retArr := make([]int, 2)
    for i,v := range nums{
        if idx, ok := m[target - v]; ok{
            retArr[0] = idx
            retArr[1] = i
            break
        }else{
            m[v] = i
        }
    }
    return retArr
}