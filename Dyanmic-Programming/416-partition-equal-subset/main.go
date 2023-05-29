// https://leetcode.com/problems/partition-equal-subset-sum/

package main

import "fmt"

func canPartition(nums []int) bool {
	total := 0
	for i, _ := range nums {
		total = total + nums[i]
	}

	// if this is even then partition is possible
	if total%2 != 0 {
		return false
	}

	total = total / 2
	dp := make([]bool, total+1)
	dp[0] = true

	for _, n := range nums {
		for i := total; i >= n; i-- {
			dp[i] = dp[i] || dp[i-n]
		}
	}
	return dp[total]
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))
}
