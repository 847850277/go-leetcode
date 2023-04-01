package march17

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	n := len(nums)
	f := make([]int, n)
	f[0] = nums[0]
	for i := 1; i < n; i++ {
		f[i] = f[i-1] + nums[i]
	}
	ans := []int{}
	for _, q := range queries {
		idx := sort.SearchInts(f, q+1)
		ans = append(ans, idx)
	}
	return ans
}
