package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	ret := ""
	if len(strs) > 0 {
		isCommon := true
		for i, v := range strs[0] {
			for _, str := range strs[1:] {
				if i < len(str) {
					if rune(str[i]) != v {
						isCommon = false
						break
					}
				} else {
					isCommon = false
					break
				}
			}
			if !isCommon {
				break
			}
			ret += string(v)
		}
	}
	return ret
}

func main() {
	// strs := []string{"flower", "flow", "flight"}
	// strs := []string{"dog", "racecar", "car"}
	strs := []string{}
	fmt.Printf("longest common prefix: %s\r\n", longestCommonPrefix(strs))
}
