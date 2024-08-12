package main

func minOperations(nums []int, k int) int {
	a := make(map[int]struct{}, k)
	cnt := 0
	for elem := range k {
		a[elem+1] = struct{}{}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		delete(a, nums[i])
		cnt++
		if len(a) == 0 {
			return cnt
		}
	}
	return -1
}
