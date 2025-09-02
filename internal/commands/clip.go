package commands

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

var clipCmd = &cobra.Command{
	Use:   "clip",
	Short: "Process text from clipboard",
	Long:  `Read text from clipboard, process it, and generate a structured Markdown note`,
	Run: func(cmd *cobra.Command, args []string) {
		text, err := clipboard.ReadAll()
		if err != nil {
			fmt.Printf("Error reading from clipboard: %v\n", err)
			return
		}

		if text == "" {
			fmt.Println("Clipboard is empty")
			return
		}

		fmt.Println("Received text from clipboard:")
		fmt.Println("------------------------")
		fmt.Println(text)
		fmt.Println("------------------------")
		fmt.Println("Text processing will be implemented in the next steps.")
	},
}

func init() {
	rootCmd.AddCommand(clipCmd)
}
