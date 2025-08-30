package main

import (
	"github.com/hi10drasingh/coding-challenge-ccwc/internal/service"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ccwc <file>",
		Short: "ccwc is a command-line tool just like wc in linux.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]

			options := service.Options{}
			countBytes, _ := cmd.Flags().GetBool("count-bytes")
			if countBytes {
				options.CountBytes = true
			}

			countLines := cmd.Flags().Changed("count-lines")
			if countLines {
				options.CountLines = true
			}

			countWords := cmd.Flags().Changed("count-words")
			if countWords {
				options.CountWords = true
			}

			countCharacters := cmd.Flags().Changed("count-characters")
			if countCharacters {
				options.CountCharacters = true
			}

			service.Process(filePath, options)
		},
	}

	rootCmd.Version = "1.0.0"

	rootCmd.Flags().BoolP("count-bytes", "c", false, "Count the number of bytes in the file")
	rootCmd.Flags().BoolP("count-lines", "l", false, "Count the number of lines in the file")
	rootCmd.Flags().BoolP("count-words", "w", false, "Count the number of words in the file")
	rootCmd.Flags().BoolP("count-characters", "m", false, "Count the number of characters in the file")

	return rootCmd
}

func main() {
	rootCmd := NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
	}
}
