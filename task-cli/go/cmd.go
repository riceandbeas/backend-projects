package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new task",
	RunE: func(md *cobra.Command, args []string) error {
		tracker, err := NewTracker()
		if err != nil {
			return fmt.Errorf("could not create tracker: %w", err)
		}

		if len(args) < 1 {
			return fmt.Errorf("a description for the task is needed")
		}

		id := tracker.addTask(args[0])
		if err := tracker.saveTasks(); err != nil {
			return fmt.Errorf("could not save task: %w", err)
		}

		fmt.Printf("Task added successfully (ID: %d)\n", id)
		return nil
	},
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates a task's description",
	RunE: func(md *cobra.Command, args []string) error {
		tracker, err := NewTracker()
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

		tracker.updateTask(id, args[1])
		if err := tracker.saveTasks(); err != nil {
			return fmt.Errorf("could not save task %w", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(updateCmd)
}

var rootCmd = &cobra.Command{
	Use:   "taskcli",
	Short: "A CLI app to track your tasks and manage your to-do list",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
