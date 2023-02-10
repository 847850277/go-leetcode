package february10

/*
*
掷骰子模拟
*/
func dieSimulator(n int, rollMax []int) int {
	mod := 1000000007
	dp := [6][15]int{}
	total := [6]int{}
	for i := 0; i < 6; i++ {
		dp[i][0] = 1
		total[i] = 1
	}
	for n > 1 {
		tmp := [6][15]int{}
		for i := 0; i < 6; i++ {
			for j := 0; j < 6; j++ {
				if i != j {
					tmp[i][0] = (tmp[i][0] + total[j]) % mod
				}
			}
			for j := 1; j < rollMax[i]; j++ {
				tmp[i][j] = dp[i][j-1]
			}
		}

		dp = tmp
		for i := 0; i < 6; i++ {
			total[i] = 0
			for j := 0; j < rollMax[i]; j++ {
				total[i] = (total[i] + dp[i][j]) % mod
			}
		}
		n--
	}
	ret := 0
	for i := 0; i < 6; i++ {
		ret = (ret + total[i]) % mod
	}
	return ret
}
