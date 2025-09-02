package commands

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var stdinCmd = &cobra.Command{
	Use:   "stdin",
	Short: "Process text from standard input",
	Long:  `Read text from standard input, process it, and generate a structured Markdown note.`,
	Run: func(cmd *cobra.Command, args []string) {
		text, err := readFromStdin()
		if err != nil {
			fmt.Printf("Error reading from stdin: %v\n", err)
			return
		}

		if text == "" {
			fmt.Println("No input received. Please provide some text.")
			return
		}

		fmt.Println("Received text from stdin:")
		fmt.Println("------------------------")
		fmt.Println(text)
		fmt.Println("------------------------")
		fmt.Println("Text processing will be implemented in the next steps.")
	},
}

func readFromStdin() (string, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return "", nil // No data available
	}

	reader := bufio.NewReader(os.Stdin)
	var builder strings.Builder

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return "", err
		}

		builder.WriteString(line)

		if err == io.EOF {
			break
		}
	}

	return builder.String(), nil
}

func init() {
	rootCmd.AddCommand(stdinCmd)
}
