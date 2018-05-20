package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"bytes"
)

func execute(bf string) string {
	cursor := 0
	var output bytes.Buffer
	arr := make([]byte, 1000)
	
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
				for j := i; j < len(bf); j++ {
					if bf[j] == byte(']') {
						i = j
						break
					} 
					if j == len(bf)-1 {
						panic("Could not found ']'!")
					}
					
				}
			}
		case ']':
			if arr[cursor] != 0 {
				for j := i; j > 0 ; j-- {
					if bf[j] == byte('[') {
						i = j
						break
					} 
					if j == 0 {
						panic("Could not found '['!")
					}
				}
			}
		}
	}
	return output.String()
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