package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsASCII(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, IsASCII(" abc1"))
	a.Equal(false, IsASCII("хекслет"))
	a.Equal(false, IsASCII("hello \U0001F970"))
}
