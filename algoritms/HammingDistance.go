package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hammingDistance(x int, y int) int {
	res := strconv.FormatInt(int64(x^y), 2)
	return strings.Count(res, "1")
}
