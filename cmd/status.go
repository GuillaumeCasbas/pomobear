package cmd

import (
	"fmt"

	"github.com/guillaumecasbas/pomobear/adapters"
	"github.com/guillaumecasbas/pomobear/usecases"
	"github.com/spf13/cobra"
)

func init() {
	statusCmd.Flags().Bool("raw", false, "only the the raw number of seconds")
	rootCmd.AddCommand(statusCmd)
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Display the pomodoro remaining time",
	Long:  `Display the active pomodoro remaining time or empty`,
	RunE: func(cmd *cobra.Command, args []string) error {
		presenter := adapters.NewCmdPresenter()
		repo, err := adapters.NewPomodoroFileRepository()
		wantRaw, err := cmd.Flags().GetBool("raw")

		if err != nil {
			return err
		}
		usecase := usecases.NewPomodoroUsecases(repo)

		remainingSeconds, err := usecase.Status()

		if err != nil {
			return err
		}

		if wantRaw {
			fmt.Println(remainingSeconds)
		} else {
			fmt.Println(presenter.DisplayStatus(remainingSeconds))
		}

		return nil
	},
}
