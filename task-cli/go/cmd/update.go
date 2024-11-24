package cmd

import (
	"fmt"
	"strconv"

	"taskcli/task"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a task's description",
	RunE: func(md *cobra.Command, args []string) error {
		tracker, err := task.NewTracker()
		if err != nil {
			return fmt.Errorf("could not create tracker: %w", err)
		}

		if len(args) < 2 {
			return fmt.Errorf("the task's ID and a new description are needed")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("could not convert id to int")
		}

		if err := tracker.UpdateTask(id, args[1]); err != nil {
			return fmt.Errorf("could not update task: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
