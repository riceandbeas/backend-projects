package cmd

import (
	"fmt"

	"taskcli/task"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds a new task",
	Example: "Single word description: add study\nMulti-word description: add \"buy strawberries\"",
	RunE: func(md *cobra.Command, args []string) error {
		tracker, err := task.NewTracker()
		if err != nil {
			return fmt.Errorf("could not create tracker: %w", err)
		}

		if len(args) < 1 {
			return fmt.Errorf("a description for the task is needed")
		}

		id, err := tracker.AddTask(args[0])
		if err != nil {
			return fmt.Errorf("could not add task: %w", err)
		}

		fmt.Printf("Task added successfully (ID: %d)\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
