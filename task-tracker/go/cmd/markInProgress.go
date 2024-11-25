package cmd

import (
	"fmt"
	"strconv"

	"taskcli/task"

	"github.com/spf13/cobra"
)

var markInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "marks the task as in progress",
	RunE: func(md *cobra.Command, args []string) error {
		tracker, err := task.NewTracker()
		if err != nil {
			return fmt.Errorf("could not create tracker: %w", err)
		}

		if len(args) < 1 {
			return fmt.Errorf("the id is needed")
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("could not convert id to int")
		}

		if err := tracker.MarkInProgress(id); err != nil {
			return fmt.Errorf("could not mark in progress: %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(markInProgressCmd)
}
