package cmd

import (
	"os"

	"github.com/OliverCardoza/traindown-cli/cmd/internal"
	"github.com/spf13/cobra"
	"github.com/traindown/traindown-go"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "traindown-cli",
	Short: `traindown-cli is a command line tool for fitness data using the Traindrain spec.`,
}

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

	rootCmd.PersistentFlags().StringP("input", "i", "", "Input file or directory")
	rootCmd.MarkFlagRequired("input")

	rootCmd.PersistentFlags().String("extension", ".traindown", "The file extension used to identify traindown files. Other files are ignored.")
}

func readInput() ([]*traindown.Session, error) {
	input, err := rootCmd.Flags().GetString("input")
	if err != nil {
		panic(err)
	}
	extension, err := rootCmd.Flags().GetString("extension")
	if err != nil {
		panic(err)
	}
	reader := internal.NewTraindownReader(extension)
	return reader.Read(input)
}
