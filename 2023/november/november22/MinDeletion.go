package main

func main() {

}

func minDeletion(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	stack := make([]int, 0)
	stack = append(stack, nums[0])
	for i := 1; i < length; i++ {
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
		return length - len(stack)
	} else {
		return length - len(stack) + 1
	}
}
