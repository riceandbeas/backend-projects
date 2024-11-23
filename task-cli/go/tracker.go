package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
)

const TaskFile string = "tasks.json"

type Tracker struct {
	Tasks  []Task
	lastId int
}

func NewTracker() (*Tracker, error) {
	tasks, err := loadTasks()
	if err != nil {
		return nil, fmt.Errorf("NewTracker(): %w", err)
	}

	lastId := 0
	if len(tasks) > 0 {
		lastId = tasks[len(tasks)-1].Id
	}

	return &Tracker{
		Tasks:  tasks,
		lastId: lastId,
	}, nil
}

func loadTasks() ([]Task, error) {
	f, err := os.Open(TaskFile)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, fmt.Errorf("loadTasks(): os.Open(): %w", err)
	}
	defer f.Close()

	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("loadTasks(): os.ReadFile(): %w", err)
	}

	tasks := make([]Task, 0)
	if err := json.Unmarshal(contents, &tasks); err != nil {
		return nil, fmt.Errorf("loadTasks(): json.Unmarshal(): %w", err)
	}

	return tasks, nil
}

func (t *Tracker) saveTasks() error {
	contents, err := json.MarshalIndent(t.Tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("saveTasks(): json.MarshalIndent(): %w", err)
	}

	file, err := os.OpenFile(TaskFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return fmt.Errorf("saveTasks(): os.OpenFile(): %w", err)
	}
	defer file.Close()

	_, err = file.Write(contents)
	if err != nil {
		return fmt.Errorf("saveTasks(): file.Write(): %w", err)
	}

	return nil
}

func (t *Tracker) addTask(desc string) int {
	task := NewTask(t.lastId+1, desc)
	t.Tasks = append(t.Tasks, task)
	return task.Id
}

func (t *Tracker) updateTask(id int, desc string) {
	for i := range t.Tasks {
		if t.Tasks[i].Id == id {
			t.Tasks[i].Description = desc
			return
		}
	}

	fmt.Println("Task with ID", id, "not found")
}
