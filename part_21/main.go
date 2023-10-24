package main

import (
	"fmt"
)

func ShiftASCII(s string, step int) string {
	if step == 0 {
		return s
	}
	rString := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		char := s[i]
		rString[i] = char + byte(step)
	}
	return string(rString)
}

func main() {
	fmt.Print(ShiftASCII("bcd2", -1))

}
