package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	str := countAndSay(n - 1)
	ret := ""
	if len(str) > 0 {
		str = str + "a"
		count := 1
		for i := 0; i < len(str)-1; i++ {
			if str[i] != str[i+1] {
				ret += strconv.Itoa(count)
				ret += string(str[i])
				count = 0
			}
			count++
		}
	}

	return ret
}

func main() {
	n := 30
	fmt.Printf("count and say: %s\r\n", countAndSay(n))
}
