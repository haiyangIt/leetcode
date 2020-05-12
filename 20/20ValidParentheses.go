/*
20 Valid Parentheses
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.

Note that an empty string is also considered valid.

Algothim:

Use stack to push the left slice, once check the right slice, pop the stack.

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

	lenthStack := len(str)
	stack := make([]int, lenthStack, lenthStack)
	stackTopIndex := 0
	for i := 0; i < len(str); i++ {
		brakeType := IsBrake(str[i])
		if brakeType == None {
			return false
		}
		if brakeType%2 == 1 {
			stack[stackTopIndex] = brakeType
			stackTopIndex++
			//stack = append(stack, brakeType)
		} else {
			if stackTopIndex == 0 {
				return false
			} else {
				top := stack[stackTopIndex-1]
				if top == brakeType-1 {
					stackTopIndex--
					//stack = stack[:len(stack)-1]
					continue
				} else {
					return false
				}
			}
		}

	}
	if stackTopIndex == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(IsValid("["))
	fmt.Println(IsValid("[]"))
	fmt.Println(IsValid("[[[[]]]]"))
	fmt.Println(IsValid(""))
	fmt.Println(IsValid("1[]"))
	fmt.Println(IsValid("[{}][]"))
}
