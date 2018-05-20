package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestBrainfuck(t *testing.T) {
	expected := "Hello World!\n"
	bfCode := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	actual := execute(bfCode)
	require.Equal(t, expected, actual)
}

