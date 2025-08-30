package main

import (
	"fmt"
	"io"
	"os"

	"github.com/hi10drasingh/coding-challenge-ccwc/internal/service"
	"github.com/spf13/cobra"
)

// parseOptions reads CLI flags into service.Options
func parseOptions(cmd *cobra.Command) service.Options {
	getBool := func(name string) bool {
		val, _ := cmd.Flags().GetBool(name)
		return val
	}

	return service.Options{
		CountBytes:      getBool("count-bytes"),
		CountLines:      getBool("count-lines"),
		CountWords:      getBool("count-words"),
		CountCharacters: getBool("count-characters"),
	}
}

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "ccwc [file]",
		Short:   "ccwc is a command-line tool similar to wc in Linux.",
		Version: "1.0.0",
		Args:    cobra.MaximumNArgs(1), // 0 (pipe) or 1 (file) arg allowed
		RunE: func(cmd *cobra.Command, args []string) error {
			options := parseOptions(cmd)
			var input []byte
			var err error

			// Case 1: File provided as argument
			if len(args) > 0 {
				input, err = os.ReadFile(args[0])
				if err != nil {
					return fmt.Errorf("failed to read file: %w", err)
				}
			} else {
				// Case 2: Check if input is piped
				stat, _ := os.Stdin.Stat()
				if (stat.Mode() & os.ModeCharDevice) == 0 {
					// stdin has data
					input, err = io.ReadAll(os.Stdin)
					if err != nil {
						return fmt.Errorf("failed to read stdin: %w", err)
					}
				} else {
					// No file and no pipe
					return fmt.Errorf("no input provided (need file or piped data)")
				}
			}

			// Process file/pipe contents
			result := service.Process(input, options)

			// If no flags, print everything (like wc)
			if cmd.Flags().NFlag() == 0 {
				fmt.Printf("%d %d %d\n", result.Lines, result.Words, result.Bytes)
				return nil
			}

			// Otherwise print only requested counts
			if options.CountBytes {
				fmt.Println(result.Bytes)
				return nil
			}
			if options.CountLines {
				fmt.Println(result.Lines)
				return nil
			}
			if options.CountWords {
				fmt.Println(result.Words)
				return nil
			}
			if options.CountCharacters {
				fmt.Println(result.Characters)
				return nil
			}

			return nil
		},
	}

	// Define flags
	rootCmd.Flags().BoolP("count-bytes", "c", false, "Count the number of bytes in the file")
	rootCmd.Flags().BoolP("count-lines", "l", false, "Count the number of lines in the file")
	rootCmd.Flags().BoolP("count-words", "w", false, "Count the number of words in the file")
	rootCmd.Flags().BoolP("count-characters", "m", false, "Count the number of characters in the file")

	return rootCmd
}

func main() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
