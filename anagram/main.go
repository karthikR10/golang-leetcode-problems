package main

import (
	"fmt"
	"strings"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	chars := [26]int{}

	for i := range s {
		chars[int(s[i]-'a')]++
		chars[int(t[i]-'a')]--
	}

	for i := 0; i < 26; i++ {
		if chars[i] != 0 {
			return false
		}
	}

	return true

}

func isAnagramWithSort(s string, t string) bool {
	first := mergeSort(strings.Split("anagram", ""))
	second := mergeSort(strings.Split("nagaram", ""))

	sortedFirst := strings.Join(first, "")
	sortedSecond := strings.Join(second, "")

	if sortedFirst != sortedSecond {
		return false
	}

	return true
}

func mergeSort(data []string) []string {
	if len(data) < 2 {
		return data
	}

	first := mergeSort(data[:len(data)/2])
	second := mergeSort(data[len(data)/2:])

	return merge(first, second)
}

func merge(first []string, second []string) []string {
	final := []string{}

	i := 0
	j := 0

	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		final = append(final, first[i])
	}

	for ; j < len(second); j++ {
		final = append(final, second[j])
	}

	return final
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagramWithSort("anagram", "nagaram"))
}
