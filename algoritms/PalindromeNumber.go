package main

import (
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

