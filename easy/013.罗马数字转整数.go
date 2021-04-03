package main

import (
	"fmt"
	"strings"
)

//用于存储判断数字的结构体
//描述当前数字和下一个数字的关系，以判断是否要改变计算方式
type Roman struct {
	val  int
	next string
}

//
func constructor() map[string]Roman {
	rMap := make(map[string]Roman)
	rMap["I"] = Roman{
		val:  1,
		next: "VX",
	}
	rMap["V"] = Roman{
		val: 5,
	}
	rMap["X"] = Roman{
		val:  10,
		next: "LC",
	}
	rMap["L"] = Roman{
		val: 50,
	}
	rMap["C"] = Roman{
		val:  100,
		next: "DM",
	}
	rMap["D"] = Roman{
		val: 500,
	}
	rMap["M"] = Roman{
		val: 1000,
	}
	return rMap
}

func romanToInt(s string) int {
	rMap := constructor()
	sum := 0
	for i := 0; i < len(s); {
		curr := rMap[string(s[i])]
		if i < len(s)-1 {
			if len(curr.next) > 0 {
				//判断当前字符的下一个是否为需要求差的字符
				if strings.Index(curr.next, string(s[i+1])) >= 0 {
					sum += rMap[string(s[i+1])].val - curr.val
					i += 2
					continue
				}
			}
		}
		sum += curr.val
		i++
	}
	return sum
}

func main() {
	// s := "III"
	// s := "IV"
	// s := "IX"
	// s := "LVIII"
	s := "MCMXCIV"
	fmt.Printf("roman to int: %d\r\n", romanToInt(s))
}
