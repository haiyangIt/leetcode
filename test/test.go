package main

import "fmt"

func main() {
	testStr := "abc中文"

	for i, ch := range testStr {
		fmt.Println(fmt.Sprintf("%d:%c", i, ch))
	}
	fmt.Printf("%v\n", testStr)
	fmt.Printf("%s len:%d\n, ", testStr, len([]rune(testStr)))

	for i, ch := range []rune(testStr) {
		fmt.Println(fmt.Sprintf("%d:%c", i, ch))
	}
}
