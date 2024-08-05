package brainfuck

import (
	"fmt"
	"os"
)

func Interpret(program string, buffer []uint8) {

	lenProgram := uint32(len(program))
	programPointer := uint32(0)
	bufferPointer := uint32(0)

	for programPointer < lenProgram {
		switch program[programPointer] {

		case '>':
			bufferPointer++

		case '<':
			bufferPointer--

		case '+':
			buffer[bufferPointer]++

		case '-':
			buffer[bufferPointer]--

		case '.':
			fmt.Printf("%c", buffer[bufferPointer])

		case ',':
			b := make([]byte, 1)
			fmt.Printf("Input required enter a char: ")
			_, err := os.Stdin.Read(b)

			if err != nil {
				fmt.Println(err)
				panic("Error occured when reading input. Quitting...")
			}

			buffer[bufferPointer] = uint8(b[0])

		case '[':
			if buffer[bufferPointer] == 0 {
				// Jump past the matching ]
				copyOfProgramPointer := programPointer
				openBracketCounter := 1

			closeBracketLoop:
				for {
					programPointer++

					if programPointer >= lenProgram {
						fmt.Printf("No closing bracket found for [ at position %d. ", copyOfProgramPointer)
						panic("Quitting...")
					}

					switch program[programPointer] {
					case '[':
						openBracketCounter++
					case ']':
						openBracketCounter--

						if openBracketCounter == 0 {
							break closeBracketLoop
						}
					}
				}
			}
		case ']':
			if buffer[bufferPointer] != 0 {
				// Jump back to matching [

				copyOfProgramPointer := programPointer
				closeBracketCounter := 1

			openBracketLoop:
				for {
					if programPointer == 0 {
						fmt.Printf("No opening bracket found for ] at position %d. ", copyOfProgramPointer)
						panic("Quitting...")
					}

					programPointer--

					switch program[programPointer] {
					case ']':
						closeBracketCounter++
					case '[':
						closeBracketCounter--

						if closeBracketCounter == 0 {
							break openBracketLoop
						}
					}
				}
			}
		}

		programPointer++
	}
}
