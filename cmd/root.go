package cmd

import (
	"os"

	"github.com/spf13/cobra"
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
}
