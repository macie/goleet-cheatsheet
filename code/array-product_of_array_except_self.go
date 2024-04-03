package main

func productExceptSelf(nums []int) []int {
	size := len(nums)
	if size <= 1 {
		return nums
	}
	answer := make([]int, 0, size)

	answer = append(answer, 1)
	for i := 1; i < size; i++ {
		answer = append(answer, nums[i-1]*answer[i-1])
	}

	right := 1
	for i := size - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}

	return answer
}
