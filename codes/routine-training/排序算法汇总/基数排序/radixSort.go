package main

import "fmt"

// https://www.runoob.com/w3cnote/radix-sort.html
func radixSort(nums []int) []int {
	largestNum := findLagrestNum(nums)
	size := len(nums)
	significantDigit := 1
	semiSorted := make([]int, size)
	// Loop until we reach the largest significant digit
	for largestNum/significantDigit > 0 {
		bucket := [10]int{0}
		// Counts the number of "keys" or digits that will go into each bucket
		for i := 0; i < size; i++ {
			bucket[(nums[i]/significantDigit)%10]++
		}
		// Add the count of the previous buckets,
		// Acquires the indexes after the end of each bucket location in the array
		// Works similar to the count sort algorithm
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i-1]
		}
		// Use the bucket to fill a "semiSorted" array
		for i := size - 1; i >= 0; i-- {
			bucket[(nums[i]/significantDigit)%10]--
			semiSorted[bucket[(nums[i]/significantDigit)%10]] = nums[i]
		}
		// Replace the current array with the semisorted array
		for i := 0; i < size; i++ {
			nums[i] = semiSorted[i]
		}
		// Move to next significant digit
		significantDigit *= 10
	}
	return nums
}

func findLagrestNum(nums []int) int {
	lagrestNum := nums[0]
	for _, v := range nums {
		if v > lagrestNum {
			lagrestNum = v
		}
	}
	return lagrestNum
}

func main() {
	nums := []int{1, 4, 2, 7, 9, 6, 5, 8, 3, 10}
	fmt.Println(radixSort(nums))
}
