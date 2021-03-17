package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) <= 0 || len(nums2) <= 0 {
		return []int{}
	}
	arrDict := map[int]int{}
	for _, value := range nums1 {
		arrDict[value]++
	}
	index := 0
	for _, value := range nums2 {
		if arrDict[value] > 0 {
			arrDict[value]--
			nums2[index] = value
			index++
		}
	}
	return nums2[0:index]
}
func intersectForSortedArray(nums1 []int, nums2 []int) []int {
	if len(nums1) <= 0 || len(nums2) <= 0 {
		return []int{}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)
	fmt.Println(nums1, nums2)
	i, j, index := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			nums1[index] = nums1[i]
			index++
			i++
			j++
		}
	}
	return nums1[0:index]
}
func main() {
	nums1 := []int{4, 9, 5, 4, 4}
	nums2 := []int{9, 4, 9, 8, 4, 6}
	fmt.Println(intersect(nums1, nums2))
	fmt.Println(intersectForSortedArray(nums1, nums2))
}
