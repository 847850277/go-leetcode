package february16

func numberOfPairs(nums []int) []int {
	count := map[int]bool{}
	res := 0
	for _, num := range nums {
		count[num] = !count[num]
		if !count[num] {
			res++
		}
	}
	return []int{res, len(nums) - 2*res}
}
