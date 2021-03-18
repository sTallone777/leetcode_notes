package main

import "fmt"

func plusOne(digits []int) []int {
	len := len(digits)
	d := 0
	p := 1
	for i := len - 1; i >= 0; i-- {
		if i != len-1 {
			p = 0
		}
		if digits[i]+p+d == 10 {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			} else {
				d = 1
			}
		} else {
			digits[i] = digits[i] + p + d
			break
		}
	}
	return digits
}

func main() {
	digits := []int{1, 9}
	fmt.Printf("plus one: %v\r\n", plusOne(digits))
}
