// https://leetcode.com/problems/longest-palindromic-substring/

// 5. Longest Palindromic Substring
// Medium
// 25.2K
// 1.5K
// Companies
// Given a string s, return the longest
// palindromic

// substring
//  in s.

// Example 1:

// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
// Example 2:

// Input: s = "cbbd"
// Output: "bb"

// Constraints:

// 1 <= s.length <= 1000
// s consist of only digits and English letters.

package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + (maxLen / 2)
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	l, r := left, right

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	pal := "cbbd"
	fmt.Println(longestPalindrome(pal))
}
