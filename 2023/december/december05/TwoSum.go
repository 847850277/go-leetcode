package main

import "fmt"

func main() {
	//array := []int{2, 7, 11, 15}
	array := []int{3, 2, 4}
	//target := 9
	target := 6
	fmt.Println(twoSum(array, target))

}

/*
*
两数之和
*/
func twoSum(array []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(array); i++ {
		i2 := target - array[i]
		_, ok := m[i2]
		if ok {
			ans := []int{m[i2], i}
			return ans
		}
		m[array[i]] = i
	}
	return nil
}
