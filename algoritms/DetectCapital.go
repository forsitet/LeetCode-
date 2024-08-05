package main

import (
	"fmt"
	"strings"
)

func detectCapitalUse(word string) bool {
	if word == strings.ToLower(word)||word == strings.ToUpper(word) || word == strings.Title(strings.ToLower(word)) {
		return true
	}
	return false
}

func main() {
	fmt.Println(detectCapitalUse("Ml"))
}
