package main

func search(nums []int, target int) int {
	leftIndex:=binarySearch(nums,target)
	if leftIndex==len(nums) || nums[leftIndex]!=target{
		return 0
	}
	rightIndex:=binarySearch(nums,target+1)
	return rightIndex-leftIndex

}

func binarySearch(nums []int, target int) int {
	left,right:=0,len(nums)-1
	for left<=right{
		mid:=left+(right-left)/2
		if nums[mid]<target{
			left=mid+1
		} else {
			right=mid-1
		}
	}
	return left
}