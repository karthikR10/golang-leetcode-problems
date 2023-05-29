// https://leetcode.com/problems/longest-increasing-subsequence/

package main

import "fmt"

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	maxLength := 0
	for x := range dp {
		maxLength = max(maxLength, dp[x])
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
