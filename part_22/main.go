package main

import (
	"fmt"
	"unicode"
)

// BEGIN (write your solution here)
func IsASCII(s string) bool {
	for _, c := range s {
		if c > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func main() {
	fmt.Print(IsASCII("хекслет"))

}

// END
