// You have an array of floating point numbers averages which is initially empty. You are given an array nums of n integers where n is even.

// You repeat the following procedure n / 2 times:

// Remove the smallest element, minElement, and the largest element maxElement, from nums.
// Add (minElement + maxElement) / 2 to averages.
// Return the minimum element in averages.

package main

import "fmt"

func quick_sort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	key_elem := nums[0]
	less := []int{}
	gress := []int{}

	for _, val := range nums[1:] {
		if val < key_elem {
			less = append(less, val)
		} else {
			gress = append(gress, val)
		}
	}
	res := append(quick_sort(less), key_elem)
	res = append(res, quick_sort(gress)...)
	return res

}

func minElement(nums []float64) float64 {
	min := nums[0]
	for _, elem := range nums[1:] {
		if elem < min {
			min = elem
		}
	}
	return min
}

func minimumAverage(nums []int) float64 {
	nums = quick_sort(nums)
	n := len(nums)
	avrg := make([]float64, n/2)
	for i := 0; i < n/2; i++ {
		avrg[i] = float64(float64((nums[i] + nums[n-i-1])) / 2)
	}
	fmt.Println(avrg)
	return minElement(avrg)

}

