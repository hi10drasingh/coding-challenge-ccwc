package service

import (
	"unicode"
)

type Options struct {
	CountBytes      bool
	CountLines      bool
	CountWords      bool
	CountCharacters bool
}

type Results struct {
	Bytes      int
	Lines      int
	Words      int
	Characters int
}

func Process(bytes []byte, options Options) Results {
	input := string(bytes)
	switch {
	case options.CountBytes:
		return Results{Bytes: CountBytes(input)}

	case options.CountLines:
		return Results{Lines: CountLines(input)}

	case options.CountWords:
		return Results{Words: CountWords(input)}

	case options.CountCharacters:
		return Results{Characters: CountCharacters(input)}

	default:
		return Results{
			Bytes:      CountBytes(input),
			Lines:      CountLines(input),
			Words:      CountWords(input),
			Characters: CountCharacters(input),
		}
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
