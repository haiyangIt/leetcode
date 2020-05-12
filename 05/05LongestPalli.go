/*
Longest Palindromic Substring

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"

对于第j项，判断
 1.  j == j+1
	 j- 1 == j + 2
	 j - 2 == j + 3
 2. j- 1 == j+1

*/

package main

import "fmt"

func isPalindrome(s string, index int) string {

	subStr1 := s[0:1]

	j := 1
	lenStr := len(s)

	for true {
		if index+j > lenStr-1 || index-j+1 < 0 {
			break
		}
		if s[index-j+1] != s[index+j] {
			break
		}
		subStr1 = s[index-j+1 : index+j+1]
		j++
	}

	j = 1
	subStr2 := ""
	for true {
		if index-j < 0 || index+j > lenStr-1 {
			break
		}

		if s[index-j] != s[index+j] {
			break
		}
		subStr2 = s[index-j : index+j+1]
		j++
	}
	if len(subStr1) > len(subStr2) {
		return subStr1
	} else {
		return subStr2
	}
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	maxStr := ""
	for index := range []byte(s) {
		tempStr := isPalindrome(s, index)
		if len(maxStr) < len(tempStr) {
			maxStr = tempStr
		}
	}
	return maxStr
}

func main() {
	testStr := "ccc"
	fmt.Println(fmt.Sprintf("%s 's longest string is :%s \n", testStr, longestPalindrome(testStr)))

	testStr = "a"
	fmt.Println(fmt.Sprintf("%s 's longest string is :%s \n", testStr, longestPalindrome(testStr)))

	testStr = "ac"
	fmt.Println(fmt.Sprintf("%s 's longest string is :%s \n", testStr, longestPalindrome(testStr)))

	testStr = "abccbaa"
	fmt.Println(fmt.Sprintf("%s 's longest string is :%s \n", testStr, longestPalindrome(testStr)))

	testStr = "abcbabcba"
	fmt.Println(fmt.Sprintf("%s 's longest string is :%s \n", testStr, longestPalindrome(testStr)))
}
