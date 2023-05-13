package may13

import "sort"

func findMaxK(nums []int) int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	for left <= right {
		sum := nums[left] + nums[right]
		if sum > 0 {
			right--
		} else if sum < 0 {
			left++
		} else {
			return nums[right]
		}
	}
	return -1
}
