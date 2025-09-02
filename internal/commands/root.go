package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "brain",
	Short: "Second Brain Automatization - Obsidian note generator",
	Long: `A CLI tool that helps you create enriched Markdown notes for your 
	Obsidian vault from articles you read. It process text and generates 
	structured notes with summaries, key takeaways, and additional information.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

}
