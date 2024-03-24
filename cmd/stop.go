
package cmd

import (
	"fmt"

	"github.com/guillaumecasbas/pomobear/adapters"
	"github.com/guillaumecasbas/pomobear/usecases"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a pomodoro",
	Long:  `Stop a pomodoro`,
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := adapters.NewPomodoroFileRepository()
		if err != nil {
            return err
		}
		usecase := usecases.NewPomodoroUsecases(repo)

        ok, err := usecase.Stop()

		if err != nil {
            return err
		}

        if !ok {
            fmt.Println("None pomodoro is running")
        }

        return nil
	},
}
