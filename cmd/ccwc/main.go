package main

import (
	"github.com/hi10drasingh/coding-challenge-ccwc/internal/service"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "ccwc <file>",
		Short: "ccwc is a command-line tool just like wc in linux.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]
			options := make(map[string]bool)

			countBytes, _ := cmd.Flags().GetBool("count-bytes")
			if countBytes {
				options["-c"] = true
			}

			service.Process(filePath, options)
		},
	}

	rootCmd.Version = "1.0.0"

	rootCmd.Flags().BoolP("count-bytes", "c", false, "Count the number of bytes in the file")

	if err := rootCmd.Execute(); err != nil {
		println(err.Error())
	}
}
