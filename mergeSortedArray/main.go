package main

// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
// Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
// The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	for k := m + n - 1; k >= 0; k-- {
		if (i >= 0 && j >= 0 && nums1[i] > nums2[j]) || j < 0 {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}

	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	m := 3
	n := 3
	merge(nums1, m, nums2, n)
}
