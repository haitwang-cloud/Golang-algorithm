package BubbleSort

func BubbleSortPro(nums []int) []int {
	length := len(nums)
	var flag bool
	for i := 0; i < length && flag; i++ {
		for j := 0; j < length-i-1; j++ {
			//如果没发生交换，则依旧为false，下次就会跳出循环
			flag = false
			if nums[j] > nums[j+1] {
				//发生交换，则变为true,下次继续判断
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
	}
	return nums
}

func BubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}
