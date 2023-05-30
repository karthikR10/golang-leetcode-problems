// https://leetcode.com/problems/unique-paths/

package main

import "fmt"

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[m-1][n-1] = 1

	for i := 0; i < m-1; i++ {
		dp[i][n-1] = 1
	}

	for j := 0; j < n-1; j++ {
		dp[m-1][j] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i][j+1] + dp[i+1][j]
		}
	}

	return dp[0][0]
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}

// m = 2 n =3
// dp[0][1] = dp[0][2] + dp[1][1]
// [] [] [1]
// [1] [1] [1]

// i = 0
// j = 1
