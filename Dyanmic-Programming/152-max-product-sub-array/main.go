package main

import "fmt"

func maxProduct(nums []int) int {
	res := nums[0]
	curMax, curMin := 1, 1
	for i := 0; i < len(nums); i++ {
		temp := curMax * nums[i]
		curMax = max(max(curMax*nums[i], curMin*nums[i]), nums[i])
		curMin = min(min(temp, curMin*nums[i]), nums[i])
		res = max(res, curMax)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(nums))
}
