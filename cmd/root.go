/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "My version of grep",
	Long:  "command-line utility for searching plain-text data sets for lines that match a regular expression.\ng/re/p (global / regular expression search / and print)",
	Run:   grep,
}

func init() {
	grepCmd.Flags().BoolP("ignore-case", "i", false, "Ignores case distinctions in patterns and input data")
	grepCmd.Flags().BoolP("invert-match", "v", false, "Selects the non-matching lines of the provided input pattern")
	grepCmd.Flags().BoolP("line-number", "n", false, "Prefix each line of the matching output with the line number in the input file")
	grepCmd.Flags().BoolP("word-regexp", "w", false, "Find the exact matching word from the input file or string")
	grepCmd.Flags().BoolP("count", "c", false, "Count the number of occurrences of the provided pattern instead of outputing the lines")
}

func grep(cmd *cobra.Command, args []string) {

}

func Execute() {
	err := grepCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
