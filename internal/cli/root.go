package cli

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "mws",
	Short: "Profile Manager",
	Long:  "CLI tool for managing YAML profiles",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(profileCmd)
	rootCmd.AddCommand(helpCmd)
}
