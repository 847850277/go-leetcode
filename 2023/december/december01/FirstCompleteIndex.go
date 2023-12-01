package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 2}
	mat := [][]int{{1, 4}, {2, 3}}
	fmt.Println(firstCompleteIndex(arr, mat))

}

func firstCompleteIndex(arr []int, mat [][]int) int {
	n := len(mat)
	m := len(mat[0])
	myMap := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			myMap[mat[i][j]] = []int{i, j}
		}
	}
	rowCnt := make([]int, n)
	colCnt := make([]int, m)
	for i := 0; i < n; i++ {
		rowCnt[i] = 0
	}
	for i := 0; i < m; i++ {
		colCnt[i] = 0
	}
	for i := 0; i < len(arr); i++ {
		v := myMap[arr[i]]
		rowCnt[v[0]]++
		if rowCnt[v[0]] == m {
			return i
		}
		colCnt[v[1]]++
		if colCnt[v[1]] == n {
			return i
		}
	}
	return -1
}
