package main

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for _, num := range nums[1:] {
		if sum < 0 {
			sum = num
		} else {
			sum += num
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
