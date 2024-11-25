package task

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strconv"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

func (t *Tracker) ListTasks(status string) {
	if len(status) > 0 {
		for i := range t.Tasks {
			if t.Tasks[i].Status == NewStatus(status) {
				t.printTask(t.Tasks[i])
			}
		}
	} else {
		t.printTasks(t.Tasks)
	}
}

func (t *Tracker) AddTask(desc string) (int, error) {
	task := NewTask(t.lastId+1, desc)
	t.Tasks = append(t.Tasks, task)

	if err := t.saveTasks(); err != nil {
		return 0, fmt.Errorf("AddTask(): %w", err)
	}

	return task.Id, nil
}

func (t *Tracker) UpdateTask(id int, desc string) error {
	task, err := t.findTask(id)
	if err != nil {
		return err
	}

	task.Description = desc
	task.UpdatedAt = time.Now()

	if err := t.saveTasks(); err != nil {
		return fmt.Errorf("UpdateTask(): %w", err)
	}

	return nil
}

func (t *Tracker) MarkInProgress(id int) error {
	task, err := t.findTask(id)
	if err != nil {
		return err
	}

	task.Status = InProgress
	task.UpdatedAt = time.Now()

	if err := t.saveTasks(); err != nil {
		return fmt.Errorf("MarkInProgress(): %w", err)
	}

	return nil
}

func (t *Tracker) MarkDone(id int) error {
	task, err := t.findTask(id)
	if err != nil {
		return err
	}

	task.Status = Done
	task.UpdatedAt = time.Now()

	if err := t.saveTasks(); err != nil {
		return fmt.Errorf("MarkDone(): %w", err)
	}

	return nil
}

func (t *Tracker) DeleteTask(id int) error {
	index := -1
	for i, task := range t.Tasks {
		if task.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("task not found")
	}

	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)

	if err := t.saveTasks(); err != nil {
		return fmt.Errorf("DeleteTask(): %w", err)
	}

	return nil
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

func (t *Tracker) findTask(id int) (*Task, error) {
	for i := range t.Tasks {
		if t.Tasks[i].Id == id {
			return &t.Tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}

func (t *Tracker) printTask(task Task) {
	baseStyle := lipgloss.NewStyle().
		Bold(true)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Width(40).
		Height(2).
		Underline(true).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("212"))

	border := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Center).
		BorderForeground(lipgloss.Color("63")).
		Foreground(lipgloss.Color("212"))

	fmt.Println(
		border.Render(
			titleStyle.Render("Task #"+strconv.Itoa(task.Id))+"\n",
			baseStyle.Render("Description:"), task.Description+"\n",
			baseStyle.Render("Status:"), string(task.Status),
		),
	)
}

func (t *Tracker) printTasks(tasks []Task) {
	style := lipgloss.NewStyle()
	headerStyle := style.
		Foreground(lipgloss.Color("212")).
		Bold(true)

	table := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(style.Foreground(lipgloss.Color("63"))).
		Headers(
			headerStyle.Render("ID"),
			headerStyle.Render("Description"),
			headerStyle.Render("Status"),
		)

	for _, task := range tasks {
		descStyle := style
		if task.Status == Done {
			descStyle = descStyle.Strikethrough(true)
		}
		table.Row(
			style.Render(strconv.Itoa(task.Id)),
			descStyle.Render(task.Description),
			style.Render(string(task.Status)),
		)
	}

	fmt.Println(table)
}
