package task

import (
	"strings"
	"time"
)

type Task struct {
	Id          int
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Status string

const (
	ToDo       Status = "To do"
	InProgress Status = "In progress"
	Done       Status = "Done"
)

func NewStatus(s string) Status {
	formatted := strings.ToLower(s)
	formatted = strings.Replace(formatted, "-", " ", -1)
	switch formatted {
	case "in progress":
		return InProgress
	case "done":
		return Done
	default:
		return ToDo
	}
}

func NewTask(id int, desc string) Task {
	return Task{
		Id:          id,
		Description: desc,
		Status:      NewStatus(""),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
