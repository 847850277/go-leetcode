package main

import (
	"math"
)

func main() {
	s := "0011000111"
	print(findTheLongestBalancedSubstring(s))
}

func findTheLongestBalancedSubstring(s string) int {
	res := 0
	count := make([]int, 2)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count[1]++
			res = int(math.Max(float64(res), 2*math.Min(float64(count[0]), float64(count[1]))))
		} else if i == 0 || s[i-1] == '1' {
			count[0] = 1
			count[1] = 0
		} else {
			count[0]++
		}
	}
	return res
}
