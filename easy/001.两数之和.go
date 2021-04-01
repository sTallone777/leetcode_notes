//Because the value to be obtained can be understood as target-x, 
//repeated calculations can be reduced by using the hashmap. 
//The answer is when target-x is encountered in the loop.

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
