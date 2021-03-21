package main

import "fmt"

func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])

}

func reverse(array []int) {
	for i := 0; i < len(array)/2; i++ {
		array[i], array[len(array)-i-1] = array[len(array)-i-1], array[i]
	}

}
func main() {
	test, k := []int{2, 3, 1, 5, 6}, 3
	rotate(test, k)
	fmt.Println(test)
}
