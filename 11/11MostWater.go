package main

import "fmt"

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.


*/

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}

	begin := 0
	end := len(height) - 1
	maxArea := 0

	long := 0
	wide := 0

	direction := 0 // 0 means move begin, 1 means move end.

	for begin < end {
		if height[begin] < height[end] {
			direction = 0
			long = height[begin]
		} else {
			direction = 1
			long = height[end]
		}

		wide = end - begin
		area := long * wide
		if maxArea < area {
			maxArea = area
		}

		if direction == 0 {
			for begin++; begin < end && height[begin] < height[begin-1]; begin++ {
			}
		} else {
			for end--; end > begin && height[end] < height[end+1]; end-- {
			}
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{2, 5, 3, 6, 7, 4}))
}
