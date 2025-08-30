package service

import (
	"os"
	"unicode"
)

func Process(filePath string, options map[string]bool) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	input := string(data)

	switch {
	case options["-c"]:
		byteCount := CountBytes(input)
		println(byteCount)

	case options["-l"]:
		lineCount := CountLines(input)
		println(lineCount)

	case options["-w"]:
		wordCount := CountWords(input)
		println(wordCount)

	default:
		println("No valid option provided.")
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
