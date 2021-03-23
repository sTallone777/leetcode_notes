package main

import "fmt"

func generate(numRows int) [][]int {
	var retArr [][]int
	for a := 0; a < numRows; a++ {
		if a == 0 {
			retArr = append(retArr, []int{1})
		} else {
			arr := make([]int, a+1)

			for i := range arr {
				left := 0
				right := 0
				if i-1 >= 0 && i-1 < len(retArr[a-1]) {
					left = retArr[a-1][i-1]
				}
				if i >= 0 && i < len(retArr[a-1]) {
					right = retArr[a-1][i]
				}
				arr[i] = left + right
			}
			retArr = append(retArr, arr)
		}
	}
	return retArr
}

func main() {
	fmt.Printf("generate: %v\r\n", generate(10))
}
