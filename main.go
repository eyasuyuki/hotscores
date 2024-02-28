package main

import (
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "hotscores <svn repos path>",
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Long: "Calculate hotspot scores from Subversion repository",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
