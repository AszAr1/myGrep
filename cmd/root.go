package cmd

import (
	"errors"
	"fmt"
	"github.com/TwiN/go-color"
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
	ignoreCase := getFlag(cmd, "ignore-case")
	invertMatch := getFlag(cmd, "invert-match")
	lineNumber := getFlag(cmd, "line-number")
	wordRegexp := getFlag(cmd, "word-regexp")
	count := getFlag(cmd, "count")

	lineNumberVar := 1
	countVar := 0

	regex, fileNames := getRegexAndFileNames(args)
	if wordRegexp {
		regex = "\\b" + regex + "\\b"
	}
	if ignoreCase {
		regex = "(?i)" + regex
	}

	r, err := regexp.Compile(regex)
	check(err, fmt.Sprintf("Could not compile regex '%s'\n", regex))
	for _, fileName := range fileNames {
		text := strings.Split(getTextFromFile(fileName), "\n")

		for _, l := range text {
			if r.MatchString(l) && !invertMatch {
				outputLine(strings.Join(paintedLine(r, strings.Split(l, " ")), " "), lineNumber, &lineNumberVar)
			} else if !r.MatchString(l) && invertMatch {
				outputLine(l, lineNumber, &lineNumberVar)
			}
		}

	}
}

func check(err error, errString string) {
	if err != nil {
		log.Fatal(errString, err)
	}
}

func getFlag(cmd *cobra.Command, name string) bool {
	flag, err := cmd.Flags().GetBool(name)
	check(err, fmt.Sprintf("Could not get flag %s\n", name))
	return flag
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

func paintedLine(r *regexp.Regexp, line []string) []string {
	for j, w := range line {
		if r.MatchString(w) {
			match := r.FindString(line[j])
			line[j] = strings.Replace(line[j], match, color.Ize(color.Red, match), 1)
		}
	}
	return line
}

func outputLine(line string, lineCount bool, lineCountVar *int) {
	if lineCount {
		fmt.Printf("%d: ", *lineCountVar)
		*lineCountVar++
	}
	fmt.Println(line)
}

func Execute() {
	err := grepCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
