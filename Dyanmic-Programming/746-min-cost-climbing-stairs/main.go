package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	cost = append(cost, 0)

	for i := len(cost) - 3; i >= 0; i-- {
		println(i)
		cost[i] = min(cost[i]+cost[i+1], cost[i]+cost[i+2])
	}
	fmt.Println(cost)
	return min(cost[0], cost[1])
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	cost := []int{10, 15, 20}
	fmt.Println(minCostClimbingStairs(cost))
}
