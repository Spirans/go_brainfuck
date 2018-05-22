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