package main

import (
	"fmt"
	"strings"
	"unicode"
)

// BEGIN (write your solution here)
func LatinLetters(s string) string {
	sb := strings.Builder{}
	for _, c := range s {
		if unicode.Is(unicode.Latin, c) {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func main() {
	fmt.Print(LatinLetters("hello мир!!!!!"))

}
