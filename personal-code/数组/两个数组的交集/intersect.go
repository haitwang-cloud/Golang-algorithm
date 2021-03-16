package main
import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	intersection := []int{}
	for _, num := range nums2 {
		if m[num] > 0 {
			intersection = append(intersection, num)
			m[num]--
		}
	}
	return intersection
}
func main() {
	nums1 := []int{4, 9, 5, 4, 4}
	nums2 := []int{9, 4, 9, 8, 4, 6}
	fmt.Println(nums1, nums2)
	fmt.Println(intersect(nums1, nums2))
}
