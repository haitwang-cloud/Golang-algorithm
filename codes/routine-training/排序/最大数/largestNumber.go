package main

import (
	"fmt"
	"strconv"
	"strings"
)
func Parition(arr []string, left, right int) int {
	if left < right {
		pivot := arr[left]
		for left < right {
			for left < right && pivot+arr[right] >= arr[right]+pivot {
				right--
			}
			arr[left] = arr[right]
			for left < right && pivot+arr[left] <= arr[left]+pivot {
				left++
			}
			arr[right] = arr[left]
		}
		arr[left] = pivot
	}
	return left
}
func quickSort(arr []string, left, right int) {
	if left < right {
		pivot := Parition(arr, left, right)
		quickSort(arr, left, pivot-1)
		quickSort(arr, pivot+1, right)
	}
}

func largestNumber(nums []int) string {
	arr := make([]string, len(nums))
	for i, v := range nums {
		arr[i] = strconv.Itoa(v)
	}
	quickSort(arr, 0, len(arr)-1)
	result := strings.Join(arr, "")
	if result[0] == '0' {
		return "0"
	}
	return result

}

func main() {
	nums := []int{10, 2}
	fmt.Println(largestNumber(nums))
}
