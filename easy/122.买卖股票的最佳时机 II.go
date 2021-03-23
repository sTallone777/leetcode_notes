func maxProfit(prices []int) int {
    //正常思路解法
    // p := 0
    // buyDay := 0
    // saleDay := 1
    // for saleDay < len(prices){
    //     if saleDay > buyDay{
    //         if saleDay < len(prices) -1{
    //             if prices[saleDay] - prices[buyDay] > 0 && prices[saleDay + 1] - prices[saleDay] <= 0{
    //                 p += prices[saleDay] - prices[buyDay]
    //                 buyDay = saleDay + 1
    //                 saleDay = buyDay + 1
    //             }else{
    //                 if prices[saleDay] - prices[buyDay] <= 0{
    //                     buyDay++
    //                 }
    //                 saleDay++
    //             }
    //         }else{
    //             if prices[saleDay] - prices[buyDay] > 0{
    //                 p += prices[saleDay] - prices[buyDay]
    //             }
                
    //             saleDay++
    //         }
    //     }
    // }
    // return p
    //?化解法:
    //可以理解?，想要收益最大化，那?一定要?保?出比?入价格高，因?再次?入的价格如果高于上一次?出价格，那??收益一定不是最大的。
    //因此，可以?化??到比前一个价格高的价格就可以?加?算。因?1,2,3,4,5中，2-1 + 3-2 + 4-3 + 5-4 = 5-1
    p := 0
    for i:=1; i<len(prices);i++{
        if prices[i] > prices[i-1]{
            p+=prices[i]-prices[i-1]
        }
    }
    return p
}