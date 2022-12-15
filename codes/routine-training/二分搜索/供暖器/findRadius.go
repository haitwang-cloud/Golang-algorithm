package main

import (
	"fmt"
	"math"
	"sort"
)

// https://leetcode.cn/problems/heaters/
/*

无论双指针还是二分，本质都是找当前房间最近的供暖器。
这道题可以这样考虑，供暖器的位置是确定的，所以房间所选的供暖器的位置是从自身位置开始向两边展开，
供暖器越远，供暖器需要辐射的半径就会越大。我们需要尽可能的让供暖器半径变小，
所以我们应该让能覆盖当前房间的供暖器尽可能靠近房间，所以会有以下三种情况：

    当前房间有供暖器，这是我们就选当前位置为0
    右边最近：heaters[i] - house
    左边最近：house - heaters[i]

这时我们选取最近的供暖器就可以，我们无需关心选取的供暖器会因为后面的房间的覆盖而增大半径，
因为后面的房间让半径增大后，最后的结果也会随之改变，而且之前的房间依然可以被覆盖。
*/

func findRadius(houses []int, heaters []int) int {
	//找到每个房屋位置所需要的最小半径的最大值
	ans := 0
	sort.Ints(heaters)
	for _, house := range houses {
		/*

			二分查找插入位置，结果始终是右侧的插入位置，所以结果范围是[0, len]
			这里我们搜索当前房间位置+1，所以可以给当前房间供暖的只有 i 和 i - 1
			再解释一下，我们查出来的i是 house+1 的插入位置，所以i对应的供暖器可能是 house+1
			或者 大于house但离house最近 的位置
			i - 1对应的供暖器可能是 house 或者 小于house但离house最近 的位置
			所以i和i- 1相当于覆盖了所有的可能性

		*/
		j := binarySearch(heaters, house+1)
		minDis := math.MaxInt32
		// 右边最近
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		// 左边最近
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans {
			ans = minDis
		}
	}
	return ans
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left

}

func main() {

	nums, test := []int{1, 2, 3, 4}, []int{1, 4}
	fmt.Println(findRadius(nums, test))

}
