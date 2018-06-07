package main

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestBF(t *testing.T) {
	expected := "Hello World!\n"
	bfCode := "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."

	actual := execute(bfCode)
	require.Equal(t, expected, actual)
}

func TestBFNestedLoop(t *testing.T) {
	expected := ""
	bfCode := "[-----[>]>>[<<<+++>>>[-]]"

	actual := execute(bfCode)
	require.Equal(t, expected, actual)
}

func TestTokenizer(t *testing.T) {
	expected := []Token{
		{Increment, 0},
		{Decrement, 0},
		{Next, 0},
		{Back, 0},
		{Print, 0},
		{Read,0},
		{StartLoop, 7},
		{EndLoop, 6},
	}
	input := "+-><.,[]"

	actual := tokenizer(input)
	require.Equal(t, expected, actual)
}

func TestWrapAround(t *testing.T) {
	arr := make([]byte, 1)
	arr[0] += 255
	actual := arr[0] + 1
	expected := byte(0)

	require.Equal(t, expected, actual)
}