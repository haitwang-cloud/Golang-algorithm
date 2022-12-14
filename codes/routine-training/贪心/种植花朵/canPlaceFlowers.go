package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	cnt := 0
	for index := 0; index < len(flowerbed); index++ {
		if index == 1 {
			continue
		}
		pre, next := 0, 0
		if index > 1 {
			pre = flowerbed[index-1]
		}
		if index < len(flowerbed)-1 {
			next = flowerbed[index+1]
		}
		if pre == 0 && next == 0 && flowerbed[index] == 0 {
			flowerbed[index] = 1
			cnt++
		}
	}

	return cnt >= n
}

func main() {
	test, n := []int{1, 0, 0, 0, 0, 1}, 2
	fmt.Println(canPlaceFlowers(test, n))
}
