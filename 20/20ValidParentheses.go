/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.
*/

package main

import "fmt"

const (
	None        = 0
	LeftSBrake  = 1
	RightSBrake = 2
	LeftMBrake  = 3
	RightMBrake = 4
	LeftBBrake  = 5
	RightBBrake = 6
)

func IsBrake(ch byte) int {
	switch ch {
	case '(':
		return LeftSBrake
	case ')':
		return RightSBrake
	case '[':
		return LeftMBrake
	case ']':
		return RightMBrake
	case '{':
		return LeftBBrake
	case '}':
		return RightBBrake
	default:
		return None
	}
}

func IsValid(str string) bool {
	if str == "" {
		return true
	}

	leftSlice := make([]int, 0, len(str))
	for i := 0; i < len(str); i++ {
		brakeType := IsBrake(str[i])
		if brakeType == None {
			return false
		}
		if brakeType%2 == 1 {
			leftSlice = append(leftSlice, brakeType)
		} else {
			if len(leftSlice) == 0 {
				return false
			} else {
				top := leftSlice[len(leftSlice)-1]
				if top == brakeType-1 {
					leftSlice = leftSlice[:len(leftSlice)-1]
					continue
				} else {
					return false
				}
			}
		}

	}
	return true
}

func main() {
	fmt.Println(IsValid("[]"))
	fmt.Println(IsValid("[[[[]]]]"))
	fmt.Println(IsValid(""))
	fmt.Println(IsValid("1[]"))
	fmt.Println(IsValid("[{}][]"))
}
