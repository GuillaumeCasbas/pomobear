package cmd

import (
	"fmt"

	"github.com/guillaumecasbas/pomobear/adapters"
	"github.com/guillaumecasbas/pomobear/usecases"
	"github.com/spf13/cobra"
)

func init() {
	startCmd.Flags().Int("duration", 25, "duration in minute")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a pomodoro",
	Long:  `Start a pomodoro. Default duration is 25 minutes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		duration, err := cmd.Flags().GetInt("duration")
		if err != nil {
			return err
		}

		repo, err := adapters.NewPomodoroFileRepository()
		if err != nil {
			return err
		}
		usecase := usecases.NewPomodoroUsecases(repo, usecases.WithDuration(duration))

		ok, err := usecase.Start()

		if err != nil {
			return err
		}

		if !ok {
			fmt.Println("A pomodoro is already running")
		}

		return nil
	},
}
