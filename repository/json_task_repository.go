package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/Illuminateee/task-tracker.git/entity"
)

// JSONTaskRepository implements TaskRepository using JSON file storage
type JSONTaskRepository struct {
	filePath string
}

// NewJSONTaskRepository creates a new JSON task repository
func NewJSONTaskRepository(filePath string) *JSONTaskRepository {
	return &JSONTaskRepository{
		filePath: filePath,
	}
}

// Create adds a new task to the JSON file
func (r *JSONTaskRepository) Create(task *entity.Task) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}

	tasks = append(tasks, task)
	return r.saveTasks(tasks)
}

// GetByID retrieves a task by its ID
func (r *JSONTaskRepository) GetByID(id int) (*entity.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return nil, fmt.Errorf("task with ID %d not found", id)
}

// GetAll retrieves all tasks
func (r *JSONTaskRepository) GetAll() ([]*entity.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}

	// Sort tasks by ID for consistent ordering
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].ID < tasks[j].ID
	})

	return tasks, nil
}

// GetByStatus retrieves tasks filtered by status
func (r *JSONTaskRepository) GetByStatus(status entity.TaskStatus) ([]*entity.Task, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return nil, err
	}

	var filteredTasks []*entity.Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTasks = append(filteredTasks, task)
		}
	}

	// Sort tasks by ID for consistent ordering
	sort.Slice(filteredTasks, func(i, j int) bool {
		return filteredTasks[i].ID < filteredTasks[j].ID
	})

	return filteredTasks, nil
}

// Update modifies an existing task
func (r *JSONTaskRepository) Update(updatedTask *entity.Task) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			return r.saveTasks(tasks)
		}
	}

	return fmt.Errorf("task with ID %d not found", updatedTask.ID)
}

// Delete removes a task by ID
func (r *JSONTaskRepository) Delete(id int) error {
	tasks, err := r.loadTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.ID == id {
			// Remove the task from slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			return r.saveTasks(tasks)
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

// GetNextID returns the next available ID
func (r *JSONTaskRepository) GetNextID() (int, error) {
	tasks, err := r.loadTasks()
	if err != nil {
		return 0, err
	}

	if len(tasks) == 0 {
		return 1, nil
	}

	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}

	return maxID + 1, nil
}

// loadTasks loads tasks from the JSON file
func (r *JSONTaskRepository) loadTasks() ([]*entity.Task, error) {
	// Check if file exists
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		// File doesn't exist, return empty slice
		return []*entity.Task{}, nil
	}

	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tasks file: %w", err)
	}

	// Handle empty file
	if len(data) == 0 {
		return []*entity.Task{}, nil
	}

	var tasks []*entity.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
	}

	return tasks, nil
}

// saveTasks saves tasks to the JSON file
func (r *JSONTaskRepository) saveTasks(tasks []*entity.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write tasks file: %w", err)
	}

	return nil
}
