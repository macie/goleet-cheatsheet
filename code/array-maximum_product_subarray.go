package main

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	minProduct, maxProduct, result := nums[0], nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		maxProduct = max(maxProduct*num, num)
		minProduct = min(minProduct*num, num)

		result = max(result, maxProduct)
	}

	return result
}
