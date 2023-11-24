package main

import "fmt"

func main() {
	array := []int{-1, 1, 2, 3, 1}
	target := 2
	fmt.Print(countPairs(array, target))
}

func countPairs(nums []int, target int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			i2 := nums[i] + nums[j]
			if i2 < target {
				ans++
			}
		}
	}
	return ans
}
