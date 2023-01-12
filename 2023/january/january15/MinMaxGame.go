package main

func minMaxGame(array []int) int {
	i := len(array)
	if i == 1 {
		return array[0]
	}
	newNums := make([]int, i/2)
	for j := 0; j < len(newNums); j++ {
		if j%2 == 0 {
			if array[j*2] > array[j*2+1] {
				newNums[j] = array[j*2+1]
			} else {
				newNums[j] = array[j*2]
			}
		} else {
			if array[j*2] > array[j*2+1] {
				newNums[j] = array[j*2]
			} else {
				newNums[j] = array[j*2+1]
			}
		}
	}
	return minMaxGame(newNums)
}
