package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

type Jump struct {
	jpm_type string
	position uint16
}

func execute(bf string) string {
	var cursor uint16
	var output bytes.Buffer
	jump_stack := createJumpStack(bf)
	arr := make([]byte, 65535)
	
	for i := 0; i < len(bf); i++ {
		switch bf[i] {
		case '+': arr[cursor] += 1
		case '-': arr[cursor] -= 1
		case '>': cursor += 1
		case '<': cursor -= 1
		case '.': output.WriteString(string(byte(arr[cursor])))
		case ',': 
			var input byte
			_, err := fmt.Scanf("%c", &input)
			if err != nil {
				panic(err)
			}
			arr[cursor] = input

		case '[':
			if arr[cursor] == 0 {
				i = int(jump_stack[i].position)
			}
		case ']':
			if arr[cursor] != 0 {
				i = int(jump_stack[i].position)
			} 
		}
	}
	return output.String()
}

func createJumpStack(bf string) (jumps []Jump) {
	var pointer uint16
	jump_stack := make([]uint16, 0)
	for i := 0; i < len(bf); i++ {
		switch bf[i] {
		case '>', '<', '+', '-', '.', ',':
			jumps = append(jumps, Jump{"", 0})
		case '[':
			jumps = append(jumps, Jump{"[", 0})
			jump_stack = append(jump_stack, uint16(i))
		case ']':
			if len(jump_stack) == 0 {
				return nil
			}
			pointer = jump_stack[len(jump_stack)-1]
			jump_stack = jump_stack[:len(jump_stack)-1]
			jumps = append(jumps, Jump{"]", pointer})
			jumps[pointer].position = uint16(i)
		}
	
	}
	return jumps
}

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Usage: %s filename\n", args[0])
		return
	}
	filename := args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s\n", filename)
		return
	}
	str := execute(string(content))
	fmt.Print(str)
}