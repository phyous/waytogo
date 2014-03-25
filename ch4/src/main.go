package main

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	test4p6()
}

func test4p6() {
	strings := []string{
		"asSASA ddd dsjkdsjs dk",
		"asSASA ddd dsjkdsjsこん dk"}

	for _, element := range strings {
		bytes, chars := countCharacters(element)
		fmt.Printf("The number of bytes in string %s is %d\n", element, bytes)
		fmt.Printf("The number of chars in string %s is %d\n", element, chars)
	}
}
func countCharacters(s string) (int, int){
	return len(s), utf8.RuneCountInString(s)
}
