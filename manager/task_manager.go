package manager

import (
	"fmt"

	"github.com/Illuminateee/task-tracker.git/entity"
	"github.com/Illuminateee/task-tracker.git/repository"
	"github.com/Illuminateee/task-tracker.git/usecase"
)

// TaskManager coordinates task operations and manages dependencies
type TaskManager struct {
	taskUseCase *usecase.TaskUseCase
}

// NewTaskManager creates a new task manager with JSON storage
func NewTaskManager(dataFilePath string) *TaskManager {
	taskRepo := repository.NewJSONTaskRepository(dataFilePath)
	taskUseCase := usecase.NewTaskUseCase(taskRepo)

	return &TaskManager{
		taskUseCase: taskUseCase,
	}
}

// AddTask adds a new task
func (tm *TaskManager) AddTask(title, description string) (*entity.Task, error) {
	return tm.taskUseCase.CreateTask(title, description)
}

// UpdateTask updates an existing task
func (tm *TaskManager) UpdateTask(id int, title, description string) (*entity.Task, error) {
	return tm.taskUseCase.UpdateTask(id, title, description)
}

// DeleteTask deletes a task
func (tm *TaskManager) DeleteTask(id int) error {
	return tm.taskUseCase.DeleteTask(id)
}

// GetTask gets a specific task by ID
func (tm *TaskManager) GetTask(id int) (*entity.Task, error) {
	return tm.taskUseCase.GetTask(id)
}

// ListAllTasks returns all tasks
func (tm *TaskManager) ListAllTasks() ([]*entity.Task, error) {
	return tm.taskUseCase.GetAllTasks()
}

// ListDoneTasks returns all completed tasks
func (tm *TaskManager) ListDoneTasks() ([]*entity.Task, error) {
	return tm.taskUseCase.GetTasksByStatus(entity.TaskStatusDone)
}

// ListTodoTasks returns all todo tasks
func (tm *TaskManager) ListTodoTasks() ([]*entity.Task, error) {
	return tm.taskUseCase.GetTasksByStatus(entity.TaskStatusToDo)
}

// ListInProgressTasks returns all in-progress tasks
func (tm *TaskManager) ListInProgressTasks() ([]*entity.Task, error) {
	return tm.taskUseCase.GetTasksByStatus(entity.TaskStatusInProgress)
}

// ListPendingTasks returns all non-done tasks (todo + in-progress)
func (tm *TaskManager) ListPendingTasks() ([]*entity.Task, error) {
	return tm.taskUseCase.GetPendingTasks()
}

// MarkDone marks a task as completed
func (tm *TaskManager) MarkDone(id int) (*entity.Task, error) {
	return tm.taskUseCase.MarkTaskDone(id)
}

// MarkInProgress marks a task as in progress
func (tm *TaskManager) MarkInProgress(id int) (*entity.Task, error) {
	return tm.taskUseCase.MarkTaskInProgress(id)
}

// MarkTodo marks a task as todo
func (tm *TaskManager) MarkTodo(id int) (*entity.Task, error) {
	return tm.taskUseCase.MarkTaskToDo(id)
}

// UpdateTaskStatus updates task status using string value
func (tm *TaskManager) UpdateTaskStatus(id int, statusStr string) (*entity.Task, error) {
	if !entity.IsValidStatus(statusStr) {
		return nil, fmt.Errorf("invalid status '%s'. Valid statuses are: todo, in-progress, done", statusStr)
	}

	status := entity.TaskStatus(statusStr)
	return tm.taskUseCase.UpdateTaskStatus(id, status)
}
