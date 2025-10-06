package controller

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Illuminateee/task-tracker.git/entity"
	"github.com/Illuminateee/task-tracker.git/manager"
)

// CLIController handles command line interface operations
type CLIController struct {
	taskManager *manager.TaskManager
}

// NewCLIController creates a new CLI controller
func NewCLIController(dataFilePath string) *CLIController {
	return &CLIController{
		taskManager: manager.NewTaskManager(dataFilePath),
	}
}

// HandleCommand processes CLI commands and arguments
func (c *CLIController) HandleCommand(args []string) error {
	if len(args) == 0 {
		return c.showHelp()
	}

	command := strings.ToLower(args[0])
	switch command {
	case "add":
		return c.handleAdd(args[1:])
	case "update":
		return c.handleUpdate(args[1:])
	case "delete":
		return c.handleDelete(args[1:])
	case "mark-done":
		return c.handleMarkDone(args[1:])
	case "mark-in-progress":
		return c.handleMarkInProgress(args[1:])
	case "mark-todo":
		return c.handleMarkTodo(args[1:])
	case "list":
		return c.handleList(args[1:])
	case "help", "-h", "--help":
		return c.showHelp()
	default:
		return fmt.Errorf("unknown command: %s. Use 'help' to see available commands", command)
	}
}

// handleAdd processes the add command
func (c *CLIController) handleAdd(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("add command requires a title. Usage: add \"<title>\" [\"<description>\"]")
	}

	title := args[0]
	description := ""
	if len(args) > 1 {
		description = args[1]
	}

	task, err := c.taskManager.AddTask(title, description)
	if err != nil {
		return fmt.Errorf("failed to add task: %w", err)
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	c.printTask(task)
	return nil
}

// handleUpdate processes the update command
func (c *CLIController) handleUpdate(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("update command requires ID and title. Usage: update <id> \"<title>\" [\"<description>\"]")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	title := args[1]
	description := ""
	if len(args) > 2 {
		description = args[2]
	}

	task, err := c.taskManager.UpdateTask(id, title, description)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}

	fmt.Printf("Task updated successfully\n")
	c.printTask(task)
	return nil
}

// handleDelete processes the delete command
func (c *CLIController) handleDelete(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("delete command requires a task ID. Usage: delete <id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	if err := c.taskManager.DeleteTask(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	fmt.Printf("Task %d deleted successfully\n", id)
	return nil
}

// handleMarkDone processes the mark-done command
func (c *CLIController) handleMarkDone(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("mark-done command requires a task ID. Usage: mark-done <id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	task, err := c.taskManager.MarkDone(id)
	if err != nil {
		return fmt.Errorf("failed to mark task as done: %w", err)
	}

	fmt.Printf("Task marked as done\n")
	c.printTask(task)
	return nil
}

// handleMarkInProgress processes the mark-in-progress command
func (c *CLIController) handleMarkInProgress(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("mark-in-progress command requires a task ID. Usage: mark-in-progress <id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	task, err := c.taskManager.MarkInProgress(id)
	if err != nil {
		return fmt.Errorf("failed to mark task as in progress: %w", err)
	}

	fmt.Printf("Task marked as in progress\n")
	c.printTask(task)
	return nil
}

// handleMarkTodo processes the mark-todo command
func (c *CLIController) handleMarkTodo(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("mark-todo command requires a task ID. Usage: mark-todo <id>")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("invalid task ID: %s", args[0])
	}

	task, err := c.taskManager.MarkTodo(id)
	if err != nil {
		return fmt.Errorf("failed to mark task as todo: %w", err)
	}

	fmt.Printf("Task marked as todo\n")
	c.printTask(task)
	return nil
}

// handleList processes the list command
func (c *CLIController) handleList(args []string) error {
	filter := "all"
	if len(args) > 0 {
		filter = strings.ToLower(args[0])
	}

	var tasks []*entity.Task
	var err error

	switch filter {
	case "all":
		tasks, err = c.taskManager.ListAllTasks()
	case "done":
		tasks, err = c.taskManager.ListDoneTasks()
	case "todo":
		tasks, err = c.taskManager.ListTodoTasks()
	case "in-progress":
		tasks, err = c.taskManager.ListInProgressTasks()
	case "pending":
		tasks, err = c.taskManager.ListPendingTasks()
	default:
		return fmt.Errorf("invalid filter: %s. Valid filters: all, done, todo, in-progress, pending", filter)
	}

	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}

	c.printTaskList(tasks, filter)
	return nil
}

// printTask prints a single task in a formatted way
func (c *CLIController) printTask(task *entity.Task) {
	fmt.Printf("ID: %d\n", task.ID)
	fmt.Printf("Title: %s\n", task.Title)
	if task.Description != "" {
		fmt.Printf("Description: %s\n", task.Description)
	}
	fmt.Printf("Status: %s\n", task.Status)
	fmt.Printf("Created: %s\n", task.CreatedAt.Format(time.RFC3339))
	fmt.Printf("Updated: %s\n", task.UpdatedAt.Format(time.RFC3339))
}

// printTaskList prints a list of tasks
func (c *CLIController) printTaskList(tasks []*entity.Task, filter string) {
	if len(tasks) == 0 {
		fmt.Printf("No tasks found")
		if filter != "all" {
			fmt.Printf(" with filter '%s'", filter)
		}
		fmt.Println()
		return
	}

	fmt.Printf("Tasks")
	if filter != "all" {
		fmt.Printf(" (%s)", filter)
	}
	fmt.Printf(":\n\n")

	for i, task := range tasks {
		if i > 0 {
			fmt.Println("---")
		}
		c.printTask(task)
	}
}

// showHelp displays the help message
func (c *CLIController) showHelp() error {
	helpText := `Task Tracker CLI

Usage: task-tracker <command> [arguments]

Commands:
  add "<title>" ["<description>"]     Add a new task
  update <id> "<title>" ["<desc>"]    Update an existing task
  delete <id>                         Delete a task
  mark-done <id>                      Mark task as completed
  mark-in-progress <id>               Mark task as in progress
  mark-todo <id>                      Mark task as todo
  list [filter]                       List tasks (filters: all, done, todo, in-progress, pending)
  help                                Show this help message

Examples:
  task-tracker add "Buy groceries" "Milk, bread, eggs"
  task-tracker update 1 "Buy groceries" "Milk, bread, eggs, cheese"
  task-tracker mark-done 1
  task-tracker list done
  task-tracker list pending
  task-tracker delete 1

Notes:
- Tasks are stored in tasks.json file
- Default list filter is 'all'
- 'pending' filter shows both 'todo' and 'in-progress' tasks
`
	fmt.Print(helpText)
	return nil
}
