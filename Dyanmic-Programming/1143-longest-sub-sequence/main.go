// https://leetcode.com/problems/longest-common-subsequence/
package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)

	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				dp[i][j] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	t1 := "abcde"
	t2 := "ace"
	fmt.Println(longestCommonSubsequence(t1, t2))
}

// abc
// abc

// [] [] [] 0
// [] [] [] 0
// [] [] [] 0
// 0  0  0  0
// i = 2 , j =2