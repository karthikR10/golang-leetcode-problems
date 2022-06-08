package main

import (
	"fmt"
)

// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.
func isPalindrome(data int) bool {
	given := data

	if data == 0 {
		return true
	}

	if data < 0 {
		return false
	}

	final := 0
	for ; data > 0; data /= 10 {
		fmt.Println(data)
		last := data % 10
		final = final*10 + last
		fmt.Println(final)
	}
	return given == final
}

func main() {
	fmt.Println(isPalindrome(121))
}
