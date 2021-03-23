func maxProfit(prices []int) int {
    min := math.MaxInt32
    ret := 0
    for _,v := range prices{
        if v < min{
            min = v
        }else{
            if v - min > ret {
                ret = v - min
            }
        }
    } 
    return ret
}