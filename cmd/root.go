package cmd

import (
	"fmt"
	"os"

	"github.com/OliverCardoza/traindown-cli/cmd/internal"
	"github.com/spf13/cobra"
	"github.com/traindown/traindown-go"
)

const (
	inputFlagName = "input"
	inputEnvVar   = "TRAINDOWN_INPUT"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "traindown-cli",
		Short: `traindown-cli is a command line tool for fitness data using the Traindrain spec.`,
	}
	inputFlag = internal.NewEnvFallbackFlag(inputFlagName, inputEnvVar)
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Disable the `completion` command which can be used to support shell autocomplete.
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	inputUsage := fmt.Sprintf("Input file or directory. May also be set using the %s env var.", inputEnvVar)
	rootCmd.PersistentFlags().VarP(inputFlag, inputFlagName, "i", inputUsage)
	rootCmd.PersistentFlags().String("extension", ".traindown", "The file extension used to identify traindown files. Other files are ignored.")
}

func readInput() []*traindown.Session {
	input, err := inputFlag.Get()
	if err != nil {
		exit(err)
	}

	extension, err := rootCmd.Flags().GetString("extension")
	if err != nil {
		exit(err)
	}
	reader := internal.NewTraindownReader(extension)
	sessions, err := reader.Read(input)
	if err != nil {
		exit(err)
	}
	return sessions
}

func exit(err error) {
	fmt.Printf("%v\n", err)
	os.Exit(1)
}
