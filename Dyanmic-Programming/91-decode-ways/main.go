package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(s); i++ {
		if s[i-1] != '0' {
			dp[i] = dp[i] + dp[i-1]
		}
		// if last two are in the range of 1 - 26
		if in, _ := strconv.Atoi(s[i-2 : i]); in >= 10 && in <= 26 {
			dp[i] = dp[i] + dp[i-2]
		}
	}
	return dp[len(s)]
}

func main() {
	n := "11106"
	fmt.Println(numDecodings(n))
}
