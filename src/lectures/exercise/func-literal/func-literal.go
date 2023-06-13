package main

import (
	"fmt"
	"unicode"
)

type LineCallbackFn func(line string)

func computeIterator(lines []string, callback LineCallbackFn) {
	for _, value := range lines {
		callback(value)
	}

}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	letters := 0
	spaces := 0
	digits := 0
	punctuation := 0

	counter := func(line string) {
		for _, rune := range line {
			if unicode.IsLetter(rune) {
				letters += 1
			}
			if unicode.IsDigit(rune) {
				digits += 1
			}
			if unicode.IsPunct(rune) {
				punctuation += 1
			}
			if unicode.IsSpace(rune) {
				spaces += 1
			}
		}
	}

	computeIterator(lines, counter)

	fmt.Printf("Number of letters %v\n", letters)
	fmt.Printf("Number of digits %v\n", digits)
	fmt.Printf("Number of spaces %v\n", spaces)
	fmt.Printf("Number of punctual marks %v\n", punctuation)
}
