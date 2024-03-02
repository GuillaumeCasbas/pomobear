package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pomobear",
	Short: "Pomobear - The tool to be(ar) Effective",
	Long: `Pomobear is a command line tool to help being focus and effective.
It is based on the Pomodoro technique`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
