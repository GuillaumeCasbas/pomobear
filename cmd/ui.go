package cmd

import (
	"github.com/guillaumecasbas/pomobear/adapters"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(uiCmd)
}

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Start the macos status-bar app",
	Long:  `Display the pomodoro remaining time into the status-bar
    MacOS specific feature`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := adapters.NewPomodoroFileRepository()

        if err != nil {
            return err
        }

		adapters.StartUI(repo)

		return nil
	},
}
