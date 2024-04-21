package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

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
	grepCmd.Flags().BoolP("count", "c", false, "Count the number of occurrences of the provided pattern instead of outputting the lines")
}

func grep(cmd *cobra.Command, args []string) {
	regex, fileNames := getRegexAndFileNames(args)

	r, err := regexp.Compile(regex)
	check(err, fmt.Sprintf("Could not compile regex '%s'\n", regex))
	for _, fileName := range fileNames {
		text := strings.Split(getTextFromFile(fileName), "\n")

	}
}

func check(err error, errString string) {
	if err != nil {
		log.Fatal(errString, err)
	}
}

func getRegexAndFileNames(args []string) (string, []string) {
	if len(args) == 0 {
		log.Fatal("Have not been given any file names")
	}
	return args[0], args[1:]
}

func getTextFromFile(fileName string) string {
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		log.Fatal("file doesn't exist:", fileName)
	}
	bytes, err := os.ReadFile(fileName)
	check(err, fmt.Sprintf("Could not read file %s\n", fileName))
	return string(bytes)
}

func Execute() {
	err := grepCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
