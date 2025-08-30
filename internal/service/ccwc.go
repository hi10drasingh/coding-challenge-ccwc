package service

import "os"

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
	default:
		println("No valid option provided.")
	}
}

func CountBytes(input string) int {
	return len(input)
}
