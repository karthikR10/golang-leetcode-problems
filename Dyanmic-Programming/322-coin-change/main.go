package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] = min(dp[i], 1+dp[i-c])
		}
	}
	if dp[amount] != amount+1 {
		return dp[amount]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
