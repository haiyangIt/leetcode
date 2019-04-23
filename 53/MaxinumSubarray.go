/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.


sum(i+1) = max(sum(i)+i, i)
*/

package main

import "fmt"

func maxSubArray(nums []int, max *int) int {

	if len(nums) == 0 {
		*max = 0
		return 0
	}

	if len(nums) == 1 {
		*max = nums[0]
		return nums[0]
	}

	sumi_1 := maxSubArray(nums[0:len(nums)-1], max)
	datai_1 := nums[len(nums)-1]
	maxTemp := datai_1
	if sumi_1 > datai_1 {
		maxTemp = sumi_1 + datai_1
	}

	if *max < maxTemp {
		*max = maxTemp
	}
	return maxTemp
}

func main() {
	max := 0
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSubArray(nums, &max)
	fmt.Println(max)
}
