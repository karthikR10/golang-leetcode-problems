// https://leetcode.com/problems/word-break/

package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[len(s)] = true

	for i := len(s) - 1; i >= 0; i-- {
		for _, w := range wordDict {
			if (i+len(w)) <= len(s) && s[i:i+len(w)] == w {
				dp[i] = dp[i+len(w)]
			}
			if dp[i] {
				break
			}
		}
	}
	return dp[0]
}

func main() {
	s := "catsandog"
	w := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(wordBreak(s, w))
}
