package main

import "fmt"

func twoSums(data []int, target int) []int {
	result := []int{}

	numToIndex := make(map[int]int)

	for i, val := range data {
		expected := target - val

		if expectedIndex, found := numToIndex[expected]; found && expectedIndex != i {
			result = append(result, i, expectedIndex)
			break
		}

		numToIndex[val] = i
	}

	return result
}

func main() {
	data := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSums(data, target))
}
