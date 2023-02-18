package february18

import "sort"

func findSolution(customFunction func(int, int) int, z int) (ans [][]int) {
	for x := 1; x <= 1000; x++ {
		y := 1 + sort.Search(999, func(y int) bool { return customFunction(x, y+1) >= z })
		if customFunction(x, y) == z {
			ans = append(ans, []int{x, y})
		}
	}
	return
}
