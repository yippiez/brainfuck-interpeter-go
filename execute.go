package main

import (
	"fmt"
	"os"

	"github.com/yippiez/brainfuck/brainfuck"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		fmt.Printf("Missing Filename!\nUsage: %s <filename>\n", args[0])
		return
	}

	filename := args[1]
	fileContents, err := os.ReadFile(filename)

	if err != nil {
		fmt.Printf("Error when reading file %s\n", args[0])
		return
	}

	const bufferSize = 65536 // 2^16
	buffer := make([]uint8, bufferSize)

	// fmt.Println(strings.Repeat("=", 20))

	brainfuck.Interpret(string(fileContents), buffer)

	// fmt.Println("\n" + strings.Repeat("=", 20))
}
