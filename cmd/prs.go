package cmd

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/OliverCardoza/traindown-cli/cmd/internal"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/traindown/traindown-go"
)

var prsCmd = &cobra.Command{
	Use:   "prs [pattern]",
	Short: "Lists all personal records (PRs) in the input.",
	Long: `The prs command reads a file or all files in a directory and lists all personal 
records (PRs). The output can be filtered to a specific movement by regex pattern.

Examples:

  # List all PRs in a single file:
  traindown-cli prs -i ./my_workout.traindown

  # List all PRs in a directory (only checks .traindown files):
  traindown-cli prs -i ./gym_data/

	# List PRs for squat, bench press, and deadlift:
	traindown-cli prs "Squat|Bench press|Deadlift" -i ./gym_data/
`,
	Args: cobra.MaximumNArgs(1),
	Run:  prs,
}

func init() {
	rootCmd.AddCommand(prsCmd)
}

// personalRecord is a simple struct used to hold the data for an output record.
type personalRecord struct {
	dateTime       time.Time
	movementName   string
	load           float32
	unit           string
	successfulReps int
	sets           int
}

// LessThan returns true if the personal record p is less than the r.
// When the load is equal then the number of reps and sets is used to break the tie.
// TODO: Normalize data so that units do not need comparison here.
func (p *personalRecord) LessThan(r *personalRecord) bool {
	return p.load < r.load || p.successfulReps < r.successfulReps || p.sets < r.sets
}

func newPersonalRecord(session *traindown.Session, movement *traindown.Movement, performance *traindown.Performance) *personalRecord {
	return &personalRecord{
		dateTime:     session.Date,
		movementName: movement.Name,
		load:         performance.Load,
		// TODO: Normalize data to ensure this is always filled and uses a canonical unit for comparison.
		unit:           performance.Unit,
		successfulReps: performance.Reps - performance.Fails,
		// TODO: Normalize data to ensure "Squat: 225 5r 225 5r" resolves to "Squat: 225 5r 2s"
		sets: performance.Sets,
	}
}

func prs(cmd *cobra.Command, args []string) {
	pattern := ""
	if len(args) == 1 {
		pattern = args[0]
	}
	sessions, err := readInput()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	records := make(map[string]*personalRecord)
	var sortedNames internal.StringSet
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
			for _, performance := range movement.Performances {
				currentRecord := newPersonalRecord(session, movement, performance)
				pr, found := records[movement.Name]
				if !found || pr.LessThan(currentRecord) {
					sortedNames = sortedNames.Insert(movement.Name)
					records[movement.Name] = currentRecord
				}
			}
		}
	}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Movement", "Load", "Reps", "Sets", "Date"})
	for _, name := range sortedNames {
		record, _ := records[name]
		t.AppendRow([]interface{}{record.movementName, record.load, record.successfulReps, record.sets, record.dateTime})
	}
	t.Render()
}
