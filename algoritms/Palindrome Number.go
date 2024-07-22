package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	x_str := strconv.Itoa(x)
	N := len(x_str)
	for i := 0; i <= N/2; i++ {
		if x_str[i] != x_str[N-i-1] {
			return false
		}
	}
	return true
}

func main() {
	a := 1221
	fmt.Println(isPalindrome(a))
}
