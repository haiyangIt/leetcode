package main

import "fmt"

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
			 Note that the answer must be a substring, "pwke" is a subsequence and not a substring.


Algorithm
Use an array to record the index of each character start index.

记录现在字串的起始位置nowStartIndex，和现在字串的长度nowSubStringLen
同时记录当前的最大字串的起始位置maxStartIndex，和最大的字串的长度maxSubStringLen

用一个128位的数组来记录字符串的下标，数组的索引是字符，数组的值是下标值+1

遍历字符串

	查找字符是否在数组中存在
		如果不存在
			则更新数组的值为字符索引+1
			同时nowStartIndex不变，nowSubStringLen++

			判断nowSubStringLen > maxSubStringLen :
				maxStartIndex = nowStartIndex
	 			maxSubStringLen = nowSubStringLen

		如果存在  1232341
			如果数组中的索引值-1 < 当前的字串开始索引值nowStartIndex
				则更新数组的索引值
			否则
				nowStartIndex = oldIndexOfCh -1
				nowSubStrLen = nowSubStrLen - (oldIndexOfCh -1 - nowSubStringIndex + 1)
				更新数组的索引值
*/

func lengthOfLongestSubstring(s string) (maxSubStringLen int, maxSubStringIndex int, subStr string) {
	indexOfChar := make([]int, 128, 128)
	lengthOfStr := len(s)
	if lengthOfStr == 0 {
		return 0, 0, ""
	}

	nowSubStringIndex := 0
	nowSubStringLen := 0
	maxSubStringIndex = 0
	maxSubStringLen = 0
	oldIndexOfCh := 0
	for index, ch := range []byte(s) {
		oldIndexOfCh = indexOfChar[byte(ch)] - 1

		indexOfChar[byte(ch)] = index + 1

		if oldIndexOfCh < 0 || oldIndexOfCh < nowSubStringIndex {
			nowSubStringLen++

			if maxSubStringLen < nowSubStringLen {
				maxSubStringLen = nowSubStringLen
				maxSubStringIndex = nowSubStringIndex
			}
		} else {
			nowSubStringIndex = oldIndexOfCh + 1
			nowSubStringLen = index - nowSubStringIndex + 1
		}
	}
	subStr = s[maxSubStringIndex : maxSubStringIndex+maxSubStringLen]
	return maxSubStringLen, maxSubStringIndex, subStr
}

func main() {
	testStr := "abcdab"
	length, startIndex, subStr := lengthOfLongestSubstring(testStr)
	fmt.Println(fmt.Sprintf("%s length:%d, startIndex: %d, subStr: %s", testStr, length, startIndex, subStr))

	testStr = "a"
	length, startIndex, subStr = lengthOfLongestSubstring(testStr)
	fmt.Println(fmt.Sprintf("%s length:%d, startIndex: %d, subStr: %s", testStr, length, startIndex, subStr))

	testStr = "ab"
	length, startIndex, subStr = lengthOfLongestSubstring(testStr)
	fmt.Println(fmt.Sprintf("%s length:%d, startIndex: %d, subStr: %s", testStr, length, startIndex, subStr))

	testStr = "abcdeabcd"
	length, startIndex, subStr = lengthOfLongestSubstring(testStr)
	fmt.Println(fmt.Sprintf("%s length:%d, startIndex: %d, subStr: %s", testStr, length, startIndex, subStr))

	testStr = "aaabcdaa"
	length, startIndex, subStr = lengthOfLongestSubstring(testStr)
	fmt.Println(fmt.Sprintf("%s length:%d, startIndex: %d, subStr: %s", testStr, length, startIndex, subStr))

	testStr = "aabcabcdabcabcdef"
	length, startIndex, subStr = lengthOfLongestSubstring(testStr)
	fmt.Println(fmt.Sprintf("%s length:%d, startIndex: %d, subStr: %s", testStr, length, startIndex, subStr))
	testStr = "abcbcae"
	length, startIndex, subStr = lengthOfLongestSubstring(testStr)
	fmt.Println(fmt.Sprintf("%s length:%d, startIndex: %d, subStr: %s", testStr, length, startIndex, subStr))
}
