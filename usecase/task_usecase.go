package usecase

import (
	"fmt"

	"github.com/Illuminateee/task-tracker.git/entity"
	"github.com/Illuminateee/task-tracker.git/repository"
)

// TaskUseCase handles task business logic
type TaskUseCase struct {
	taskRepo repository.TaskRepository
}

// NewTaskUseCase creates a new task use case
func NewTaskUseCase(taskRepo repository.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepo: taskRepo,
	}
}

// CreateTask creates a new task
func (uc *TaskUseCase) CreateTask(title, description string) (*entity.Task, error) {
	if title == "" {
		return nil, fmt.Errorf("task title cannot be empty")
	}

	id, err := uc.taskRepo.GetNextID()
	if err != nil {
		return nil, fmt.Errorf("failed to get next ID: %w", err)
	}

	task := entity.NewTask(id, title, description)
	if err := uc.taskRepo.Create(task); err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return task, nil
}

// GetTask retrieves a task by ID
func (uc *TaskUseCase) GetTask(id int) (*entity.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task: %w", err)
	}
	return task, nil
}

// GetAllTasks retrieves all tasks
func (uc *TaskUseCase) GetAllTasks() ([]*entity.Task, error) {
	tasks, err := uc.taskRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all tasks: %w", err)
	}
	return tasks, nil
}

// GetTasksByStatus retrieves tasks filtered by status
func (uc *TaskUseCase) GetTasksByStatus(status entity.TaskStatus) ([]*entity.Task, error) {
	tasks, err := uc.taskRepo.GetByStatus(status)
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks by status: %w", err)
	}
	return tasks, nil
}

// GetPendingTasks retrieves all non-done tasks (todo + in-progress)
func (uc *TaskUseCase) GetPendingTasks() ([]*entity.Task, error) {
	todoTasks, err := uc.taskRepo.GetByStatus(entity.TaskStatusToDo)
	if err != nil {
		return nil, fmt.Errorf("failed to get todo tasks: %w", err)
	}

	inProgressTasks, err := uc.taskRepo.GetByStatus(entity.TaskStatusInProgress)
	if err != nil {
		return nil, fmt.Errorf("failed to get in-progress tasks: %w", err)
	}

	// Combine and sort
	allPending := append(todoTasks, inProgressTasks...)
	return allPending, nil
}

// UpdateTask updates an existing task
func (uc *TaskUseCase) UpdateTask(id int, title, description string) (*entity.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task for update: %w", err)
	}

	task.Update(title, description)
	if err := uc.taskRepo.Update(task); err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return task, nil
}

// UpdateTaskStatus updates the status of a task
func (uc *TaskUseCase) UpdateTaskStatus(id int, status entity.TaskStatus) (*entity.Task, error) {
	task, err := uc.taskRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task for status update: %w", err)
	}

	task.UpdateStatus(status)
	if err := uc.taskRepo.Update(task); err != nil {
		return nil, fmt.Errorf("failed to update task status: %w", err)
	}

	return task, nil
}

// DeleteTask deletes a task by ID
func (uc *TaskUseCase) DeleteTask(id int) error {
	if err := uc.taskRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}

// MarkTaskDone marks a task as done
func (uc *TaskUseCase) MarkTaskDone(id int) (*entity.Task, error) {
	return uc.UpdateTaskStatus(id, entity.TaskStatusDone)
}

// MarkTaskInProgress marks a task as in progress
func (uc *TaskUseCase) MarkTaskInProgress(id int) (*entity.Task, error) {
	return uc.UpdateTaskStatus(id, entity.TaskStatusInProgress)
}

// MarkTaskToDo marks a task as todo
func (uc *TaskUseCase) MarkTaskToDo(id int) (*entity.Task, error) {
	return uc.UpdateTaskStatus(id, entity.TaskStatusToDo)
}
