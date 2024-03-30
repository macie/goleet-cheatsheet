package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	result := make([]int, l)

	result[0] = 1
	for i := 1; i < l; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	right := 1
	for i := l - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}
