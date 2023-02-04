package february03

import "sort"

func getMaximumConsecutive(coins []int) int {
	res := 1
	sort.Ints(coins)
	for i := 0; i < len(coins); i++ {
		if coins[i] > res {
			break
		}
		res += coins[i]
	}
	return res
}
