// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-cooldown/

package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	sold, hold, rest := 0, math.MinInt32, 0
	for i := 0; i < len(prices); i++ {
		prevSold := sold
		sold = hold + prices[i]
		hold = max(hold, rest-prices[i])
		rest = max(rest, prevSold)
	}
	return max(rest, sold)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}

// [1,2,3,0,2]

// sold, hold, rest :=  0,-infi, 0

// i = 0
//  prevsold := sold // 0    , -333, 1, 2, -1
//  sold = hold + prices[i] // -1 + 1 ,, -1 + 2 // 1,, -1 + 3 // 2 , -1 + 0 (-1), 1 + 2// 3
//  hold = max(hold, rest-prices[i]) // -1, max(-1, 0-2), -1, max(-1, 0-3), -1, max(-1, 1-0), 1, max(1, 1-2) // 1
//  rest = max(rest, prevSold) // max(0,0) 0, max(0, -3333) 0, max(0, 1), 1, max(1, 0), 1, max(1, -1)// 1

//  sold, hold, rest

//  sold = hold + prices[i]
//  hold = max(hold, rest-price[i])
//  rest = max(rest, prevSold)
