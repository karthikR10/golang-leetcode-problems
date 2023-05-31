// https://leetcode.com/problems/target-sum/

package main

import (
	"fmt"
)

// func findTargetSumWays(nums []int, target int) int {
// 	total := 0
// 	for _, n := range nums {
// 		total = total + n
// 	}

// 	if total < target || (total+target)%2 == 1 {
// 		return 0
// 	}

// 	target = (total + target) / 2
// 	target = abs(target)
// 	dp := make([]int, target+1)
// 	dp[0] = 1
// 	for _, n := range nums {
// 		for i := target; i >= n; i-- {
// 			dp[i] = dp[i] + dp[i-n]
// 		}
// 	}
// 	return dp[target]
// }

func findTargetSumWaysCachingRecur(nums []int, target int) int {
	hash := make(map[string]int)
	return findSum(nums, 0, target, hash)
}

func findSum(nums []int, index int, target int, hash map[string]int) int {
	key := fmt.Sprint(index) + "*" + fmt.Sprint(target)
	if val, ok := hash[key]; ok {
		return val
	}

	if index == len(nums) {
		if target == 0 {
			return 1
		}
		return 0
	}
	res := findSum(nums, index+1, target+nums[index], hash) + findSum(nums, index+1, target-nums[index], hash)
	hash[key] = res
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	target := 3
	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(findTargetSumWaysCachingRecur(nums, target))
}
