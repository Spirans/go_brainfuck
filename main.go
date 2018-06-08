package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

type action int

const (
	Increment action = iota
	Decrement
	Next
	Back
	Print
	Read
	StartLoop
	EndLoop
)

type Token struct {
	name   action
	jumpTo int
}

func execute(bf string) string {
	var cursor byte
	var output bytes.Buffer
	tokens := tokenizer(bf)
	arr := make([]byte, len(bf))
	
	for i := 0; i < len(bf); i++ {
		switch tokens[i].name {
		case Increment: arr[cursor] += 1
		case Decrement: arr[cursor] -= 1
		case Next: cursor += 1
		case Back: cursor -= 1
		case Print: output.WriteString(string(arr[cursor]))
		case Read:
			var input byte
			_, err := fmt.Scanf("%c", &input)
			if err != nil {
				panic(err)
			}
			arr[cursor] = input

		case StartLoop:
			if arr[cursor] == 0 {
				i = int(tokens[i].jumpTo)
			}
		case EndLoop:
			if arr[cursor] != 0 {
				i = int(tokens[i].jumpTo)
			} 
		}
	}
	return output.String()
}

func tokenizer(bf string) (tokens []Token) {
	var pointer int
	jumpStack := make([]int, 0)
	for i := 0; i < len(bf); i++ {
		switch bf[i] {
		case '+':
			tokens = append(tokens, Token{Increment, 0})
		case '-':
			tokens = append(tokens, Token{Decrement, 0})
		case '>':
			tokens = append(tokens, Token{Next, 0})
		case '<':
			tokens = append(tokens, Token{Back, 0})
		case '.':
			tokens = append(tokens, Token{Print, 0})
		case ',':
			tokens = append(tokens, Token{Read, 0})
		case '[':
			tokens = append(tokens, Token{StartLoop, 0})
			jumpStack = append(jumpStack, i)
		case ']':
			if len(jumpStack) == 0 {
				return nil
			}
			pointer = jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]
			tokens = append(tokens, Token{EndLoop, pointer})
			tokens[pointer].jumpTo = i
		}
	
	}
	return tokens
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