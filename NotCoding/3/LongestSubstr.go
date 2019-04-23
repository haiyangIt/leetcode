package main

import "fmt"

/*
3. Longest Substring Without Repeating Characters
Medium

5280

281

Favorite

Share
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

*/

func FindLongestSubStrWithoutDuplicate(s string) string {
	if len(s) == 0 {
		return ""
	}

	tempSubStrHash := make(map[byte]int)
	tempSubStr := ""
	j := 0
	for i := 0; j < len(s); {
		if existI, ok := tempSubStrHash[s[j]]; ok {
			for i != existI+1 {
				if _, ok := tempSubStrHash[s[i]]; ok {
					delete(tempSubStrHash, s[i])
				}
				i++
			}
			i = existI + 1
			tempSubStrHash[s[j]] = j
			j++
		} else {
			tempSubStrHash[s[j]] = j
			temp := s[i : j+1]
			if len(tempSubStr) <= len(temp) {
				tempSubStr = temp
			}
			j++
		}
	}
	return tempSubStr
}

func main() {
	s := "12345671357902468"
	fmt.Println(FindLongestSubStrWithoutDuplicate(s))
}
