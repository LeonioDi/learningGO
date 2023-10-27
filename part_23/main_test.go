package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatinLetters(t *testing.T) {
	a := assert.New(t)
	a.Equal("abc", LatinLetters(" abc1"))
	a.Equal("", LatinLetters(" привет"))
	a.Equal("test", LatinLetters("11 ! t e    s t %)"))
	a.Equal("John", LatinLetters("John Уолтер"))
	a.Equal("word", LatinLetters("wo[r]d"))
}
