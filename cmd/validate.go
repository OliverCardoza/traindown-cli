package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
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
}

func validate(cmd *cobra.Command, args []string) {
	sessions, err := readInput()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("Validation successful: processed %d sessions\n", len(sessions))
}
