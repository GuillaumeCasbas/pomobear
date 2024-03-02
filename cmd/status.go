package cmd

import (
	"fmt"

	"github.com/guillaumecasbas/pomobear/adapters"
	"github.com/guillaumecasbas/pomobear/usecases"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the pomodoro remaining time",
	Long:  `Display the active pomodoro remaining time or empty`,
	RunE: func(cmd *cobra.Command, args []string) error {
		presenter := adapters.NewCmdPresenter()
		repo, err := adapters.NewPomodoroFileRepository()

		if err != nil {
			return err
		}
		usecase := usecases.NewPomodoroUsecases(repo)

		remainingSeconds, err := usecase.Status()

		if err != nil {
			return err
		}

		fmt.Println(presenter.DisplayStatus(remainingSeconds))

		return nil
	},
}
