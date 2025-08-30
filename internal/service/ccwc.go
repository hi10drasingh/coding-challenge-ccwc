package service

import (
	"os"
	"unicode"
)

type Options struct {
	CountBytes      bool
	CountLines      bool
	CountWords      bool
	CountCharacters bool
}

func Process(filePath string, options Options) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := string(data)

	switch {
	case options.CountBytes:
		byteCount := CountBytes(input)
		println(byteCount)

	case options.CountLines:
		lineCount := CountLines(input)
		println(lineCount)

	case options.CountWords:
		wordCount := CountWords(input)
		println(wordCount)

	case options.CountCharacters:
		charCount := CountCharacters(input)
		println(charCount)

	default:
		byteCount := CountBytes(input)
		lineCount := CountLines(input)
		wordCount := CountWords(input)
		println(lineCount, wordCount, byteCount)
	}
}

func CountBytes(input string) int {
	return len(input)
}

func CountLines(input string) int {
	count := 0
	for _, char := range input {
		if char == '\n' {
			count++
		}
	}

	// If the input is not empty and does not end with a newline, count the last line
	if len(input) > 0 && input[len(input)-1] != '\n' {
		count++
	}

	return count
}

func CountWords(input string) int {
	count := 0
	inWord := false

	for _, char := range input {
		if unicode.IsSpace(char) {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}

	return count
}

func CountCharacters(input string) int {
	return len([]rune(input))
}
