package main

import "fmt"

func getRow(rowIndex int) []int {
	preArr := []int{}
	for a := 0; a <= rowIndex; a++ {
		curArr := make([]int, a+1)
		if a == 0 {
			curArr[a] = 1
		} else {
			for b := 0; b < len(curArr); b++ {
				if b == 0 || b == len(curArr)-1 {
					curArr[b] = 1
				} else {
					curArr[b] = preArr[b-1] + preArr[b]
				}
			}
		}
		preArr = curArr
	}
	return preArr
}
func main() {
	fmt.Printf("getRow: %v\r\n", getRow(3))
}
