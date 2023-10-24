package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShiftASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal("abc", ShiftASCII("abc", 0))
	a.Equal("bcd2", ShiftASCII("abc1", 1))
	a.Equal("abc1", ShiftASCII("bcd2", -1))
	a.Equal("rs", ShiftASCII("hi", 10))
	a.Equal("abc", ShiftASCII("abc", 256))
	a.Equal("bcd", ShiftASCII("abc", -511))
}
