package february24

func minimumOperations(array []int) int {
	set := map[int]struct{}{}
	for _, x := range array {
		if x > 0 {
			set[x] = struct{}{}
		}
	}
	return len(set)
}
