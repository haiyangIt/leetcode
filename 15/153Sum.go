/*
15. 3Sum

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

package main

import "sort"

func twoSum(nums []int, value int) [][]int {
	start := 0
	lenth := len(nums)
	end := nums[lenth-1]

	result := make([][]int, 0, lenth/2)
	for start < end {
		sum := nums[start] + nums[end]
		if sum == value {
			temp := []int{nums[start], nums[end]}
			result = append(result, temp)
			start++
			end--
		} else if sum < value {
			start++
		} else {
			end--
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	result := make([][]int, 0, length/2)
	for index, item := range nums {
		target := 0 - item
		tempResult := twoSum(nums[index+1:length], target)
		if len(tempResult) > 0 {
			for _, eachItem := range tempResult {
				temp := make([]int, 0, 3)
				temp = append(temp, item)
				temp = append(temp, eachItem...)
				result = append(result, temp)
			}
		}
	}
	return result
}

func main() {

}
