func maxProfit(prices []int) int {
    //����v�H��@
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
    //?����@:
    //�ȗ���?�C�z�v���v�ő剻�C��?���v?��?�o��?�����i���C��?�Ď�?���I���i�@�ʍ�����ꎟ?�o���i�C��??���v���s���ő�I�B
    //�����C��?��??����O�꘢���i���I���i�A��?��?�Z�B��?1,2,3,4,5���C2-1 + 3-2 + 4-3 + 5-4 = 5-1
    p := 0
    for i:=1; i<len(prices);i++{
        if prices[i] > prices[i-1]{
            p+=prices[i]-prices[i-1]
        }
    }
    return p
}