package main

import (
	"fmt"
	"unicode"
)

type lineCallback func(line string)

func lineIterator(lines []string, callback lineCallback) {
	for i := 0; i < len(lines); i++ {
		callback(lines[i])
	}
}

func main() {
	lines := []string{
		"there are",
		"68 letters",
		"five digits",
		"12 spaces",
		"and 4 punctuation marks in the lines of text!",
	}

	letters := 0
	numbers := 0
	punctuation := 0
	numberOfSpaces := 0

	linefunc := func(line string) {
		for _, r := range line {
			if unicode.IsLetter(r) {
				letters += 1
			}
			if unicode.IsDigit(r) {
				numbers += 1
			}
			if unicode.IsPunct(r) {
				punctuation += 1
			}
			if unicode.IsSpace(r) {
				numberOfSpaces += 1
			}
		}
	}

	lineIterator(lines, linefunc)
	fmt.Println(letters, "letters")
	fmt.Println(numbers, "numbers")
	fmt.Println(punctuation, "punctuation")
	fmt.Println(numberOfSpaces, "numberOfSpaces")
}
