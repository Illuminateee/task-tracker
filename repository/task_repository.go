package repository

import (
	"github.com/Illuminateee/task-tracker.git/entity"
)

// TaskRepository defines the interface for task storage operations
type TaskRepository interface {
	// Create adds a new task to the repository
	Create(task *entity.Task) error

	// GetByID retrieves a task by its ID
	GetByID(id int) (*entity.Task, error)

	// GetAll retrieves all tasks
	GetAll() ([]*entity.Task, error)

	// GetByStatus retrieves tasks filtered by status
	GetByStatus(status entity.TaskStatus) ([]*entity.Task, error)

	// Update modifies an existing task
	Update(task *entity.Task) error

	// Delete removes a task by ID
	Delete(id int) error

	// GetNextID returns the next available ID
	GetNextID() (int, error)
}
