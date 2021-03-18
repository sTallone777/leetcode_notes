package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	var longNum, shortNum string
	if len(a) >= len(b) {
		longNum = a
		shortNum = b
	} else {
		longNum = b
		shortNum = a
	}
    //将较短的字符串填零，使两个字符串长度相等，便于计算
	full0 := ""
	for i := 0; i < len(longNum)-len(shortNum); i++ {
		full0 += "0"
	}
	shortNum = full0 + shortNum

	e := 0
	f := ""
	for i := len(longNum) - 1; i >= 0; i-- {
		nLong, _ := strconv.Atoi(string(longNum[i]))
		nShort, _ := strconv.Atoi(string(shortNum[i]))
		sum := nLong + nShort + e
		if sum < 2 {
			f = strconv.Itoa(sum) + f
			e = 0
		} else {
			f = strconv.Itoa(sum%2) + f
			e = 1
		}
	}
	if e == 1 {
		f = "1" + f
	}
	return f
}

func main() {
	s1 := "1010"
	s2 := "1011"
	fmt.Printf("add binary: %s\r\n", addBinary(s1, s2))
}
