// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// You can return the answer in any order.

package main

func Newmap(nums []int) map[int]int {
	mas := make(map[int]int)
	for i, val := range nums {
		mas[val] = i
	}

	return mas
}

func twoSum(nums []int, target int) []int {
	masMap := Newmap(nums)

	for inx, val := range nums {
		temp := target - val
		if val_inx, ok := masMap[temp]; ok && val_inx != inx {
			return []int{inx, val_inx}
		}
	}
	return []int{-1}
}
