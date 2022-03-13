package cmd

import (
	"fmt"
	"regexp"

	"github.com/OliverCardoza/traindown-cli/cmd/internal"
	"github.com/spf13/cobra"
)

var movementsCmd = &cobra.Command{
	Use:   "movements [pattern]",
	Short: "Lists all movements in the input.",
	Long: `The movements command reads a file or all files in a directory and lists all movement
names.

Examples:

  # List all movements in a single file:
  traindown-cli movements -i ./my_workout.traindown

  # List all movements in a directory (only checks .traindown files):
  traindown-cli movements -i ./gym_data/

	# List movements in a directory which contain "RDL" in the name:
	traindown-cli movements RDL -i ./gym_data/
`,
	Args: cobra.MaximumNArgs(1),
	Run:  movements,
}

func init() {
	rootCmd.AddCommand(movementsCmd)
}

func movements(cmd *cobra.Command, args []string) {
	pattern := ""
	if len(args) == 1 {
		pattern = args[0]
	}
	sessions := readInput()
	var sortedNames internal.StringSet
	for _, session := range sessions {
		for _, movement := range session.Movements {
			if pattern != "" {
				matched, err := regexp.MatchString(pattern, movement.Name)
				if err != nil {
					exit(err)
				}
				if !matched {
					continue
				}
			}
			sortedNames = sortedNames.Insert(movement.Name)
		}
	}
	for _, name := range sortedNames {
		fmt.Println(name)
	}
}
