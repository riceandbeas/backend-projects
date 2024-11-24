package cmd

import (
	"fmt"

	"taskcli/task"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists tasks",
	Long:  "The list command allows you to list all tasks or filter tasks by a specific status ('todo', 'in-progress' or 'done').\nIf no keyword is provided, all tasks will be listed.",
	RunE: func(md *cobra.Command, args []string) error {
		tracker, err := task.NewTracker()
		if err != nil {
			return fmt.Errorf("could not create tracker: %w", err)
		}

		if len(args) > 0 {
			tracker.ListTasks(args[0])
			return nil
		}

		tracker.ListTasks("")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
