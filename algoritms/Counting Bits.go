package main

import (
	"strconv"
	"strings"
)

func countBits(n int) []int {
	var mas []int
	for elem := range n+1 {
		mas = append(mas, strings.Count(strconv.FormatInt(int64(elem), 2), "1"))
	}
	return mas
}
