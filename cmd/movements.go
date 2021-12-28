package cmd

import (
	"fmt"
	"os"
	"regexp"
	"sort"

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

func insertSortedSet(sortedSet []string, newString string) []string {
	index := sort.SearchStrings(sortedSet, newString)
	if index < len(sortedSet) && sortedSet[index] == newString {
		// Don't need to insert if the value is already present in the set.
		return sortedSet
	}
	sortedSet = append(sortedSet, "")
	copy(sortedSet[index+1:], sortedSet[index:])
	sortedSet[index] = newString
	return sortedSet
}

func movements(cmd *cobra.Command, args []string) {
	pattern := ""
	if len(args) == 1 {
		pattern = args[0]
	}
	sessions, err := readInput()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	var sortedNames []string
	for _, session := range sessions {
		for _, movement := range session.Movements {
			if pattern != "" {
				matched, err := regexp.MatchString(pattern, movement.Name)
				if err != nil {
					fmt.Printf("%v", err)
					os.Exit(1)
				}
				if !matched {
					continue
				}
			}
			sortedNames = insertSortedSet(sortedNames, movement.Name)
		}
	}
	for _, name := range sortedNames {
		fmt.Println(name)
	}
}
