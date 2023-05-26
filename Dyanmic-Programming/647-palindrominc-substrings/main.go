package main

import "fmt"

func countSubstrings(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		res = res + countPalindrome(s, i, i)
		res = res + countPalindrome(s, i, i+1)
	}
	return res
}

func countPalindrome(s string, left, right int) int {
	l, r := left, right
	count := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		count++
		l--
		r++
	}
	return count
}

func main() {
	pal := "aaa"
	fmt.Println(countSubstrings(pal))
}
