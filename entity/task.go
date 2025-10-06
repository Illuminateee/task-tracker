package entity

import (
	"time"
)

// TaskStatus represents the status of a task
type TaskStatus string

const (
	TaskStatusToDo       TaskStatus = "todo"
	TaskStatusInProgress TaskStatus = "in-progress"
	TaskStatusDone       TaskStatus = "done"
)

// Task represents a task entity
type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

// NewTask creates a new task with default values
func NewTask(id int, title, description string) *Task {
	now := time.Now()
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      TaskStatusToDo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateStatus updates the task status and timestamp
func (t *Task) UpdateStatus(status TaskStatus) {
	t.Status = status
	t.UpdatedAt = time.Now()
}

// Update updates task fields and timestamp
func (t *Task) Update(title, description string) {
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	t.UpdatedAt = time.Now()
}

// IsValidStatus checks if the given status is valid
func IsValidStatus(status string) bool {
	switch TaskStatus(status) {
	case TaskStatusToDo, TaskStatusInProgress, TaskStatusDone:
		return true
	default:
		return false
	}
}
