package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	result := nums[0]
	sum := 0
	for _, val := range nums {
		temp := sum + val
		result = max(result, temp)
		if temp < 0 {
			sum = 0
		} else {
			sum = temp
		}
	}

	return result
	// divide and conquer
	// return findSubArray(nums, 0, len(nums)-1)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func findSubArray(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	}

	mid := left + (right-left)/2

	a := findSubArray(nums, left, mid)
	b := findSubArray(nums, mid+1, right)

	c := findCrossSubArray(nums, left, right)

	fmt.Printf("a value %v\n", a)
	fmt.Printf("b value %v\n", a)
	fmt.Printf("c value %v\n", c)
	if a > b {
		if a > c {
			return a
		} else {
			return c
		}
	} else {
		if b > c {
			return b
		} else {
			return c
		}
	}
}

func findCrossSubArray(nums []int, left int, right int) int {
	fmt.Printf("Cross sub array input %v, %v, %v\n", nums, left, right)
	mid := left + (right-left)/2
	leftSum := math.MinInt32
	rightSum := math.MinInt32
	sum := 0

	for i := left; i <= mid; i++ {
		sum += nums[i]
		if sum > leftSum {
			leftSum = sum
		}
	}

	sum = 0
	for j := mid + 1; j <= right; j++ {
		sum += nums[j]
		if sum > rightSum {
			rightSum = sum
		}
	}

	fmt.Printf("Total sum %v\n", leftSum+rightSum)
	return leftSum + rightSum
}

func main() {
	data := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	max := maxSubArray(data)
	fmt.Printf("Final max value %v\n", max)
}
