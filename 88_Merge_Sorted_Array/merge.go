package main

import "fmt"

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{0}, 0, []int{1}, 1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	z := m + n - 1

	for j >= 0 {
		if i >= 0 && (nums1[i] > nums2[j]) {
			nums1[z] = nums1[i]
			i = i - 1
		} else {
			nums1[z] = nums2[j]
			j = j - 1
		}
		z--
	}

	fmt.Println(nums1)
}
