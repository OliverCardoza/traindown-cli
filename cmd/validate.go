package cmd

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/OliverCardoza/traindown-go"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

const (
	traindownSuffix = ".traindown"
)

var (
	input       string
	validateCmd = &cobra.Command{
		Use:   "validate",
		Short: "Confirms that a file or directory contains valid traindown content.",
		Long: `The validate command reads a file or all files in a directory and confirms they contain
valid, parseable traindown content. An error is logged on the first instance of invalid syntax.

Examples:

  # Validate a single file:
  traindown-cli validate -i ./my_workout.traindown

  # Validate a directory (only checks .traindown files):
  traindown-cli validate -i ./gym_data/
`,
		Run: validate,
	}
)

func init() {
	rootCmd.AddCommand(validateCmd)

	validateCmd.Flags().StringVarP(&input, "input", "i", "", "Input file or directory")
	validateCmd.MarkFlagRequired("input")
}

func validateFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = traindown.ParseByte(data)
	if err != nil {
		return errors.Wrapf(err, "Validation error in file: %s\n", path)
	}
	return nil
}

func validate(cmd *cobra.Command, args []string) {
	log.Printf("Running validate: %s", input)

	info, err := os.Stat(input)
	if err != nil {
		log.Fatal(err)
	}

	filesProcessed := 0
	var validateErr error
	if info.IsDir() {
		validateErr = filepath.Walk(input, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				log.Fatal(err)
			}
			if strings.HasSuffix(path, traindownSuffix) {
				filesProcessed += 1
				return validateFile(path)
			}
			return nil
		})
	} else {
		filesProcessed = 1
		validateErr = validateFile(input)
	}

	if validateErr != nil {
		log.Fatal(validateErr)
	} else {
		log.Printf("Validation successful: processed %d files", filesProcessed)
	}
}
