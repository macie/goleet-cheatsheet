package main

func maxSubArray(nums []int) int {
	// Kadane's algorithm
	largestSum, sum := nums[0], nums[0]
	for _, n := range nums[1:] {
		sum = max(n, sum+n)
		largestSum = max(largestSum, sum)
	}

	return largestSum
}
