package main

func main() {
	//array := []int{1, 1, 2, 3, 5}
	array := []int{1, 1, 2, 2, 3, 3}
	print(minDeletion(array))

}

func minDeletion(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, nums[0])
	for i := 1; i < n; i++ {
		if len(stack)%2 == 0 {
			stack = append(stack, nums[i])
		} else {
			v := stack[len(stack)-1]
			if nums[i] == v {
				continue
			}
			stack = append(stack, nums[i])
		}
	}
	if len(stack)%2 == 0 {
		return n - len(stack)
	} else {
		return n - len(stack) + 1
	}
}
