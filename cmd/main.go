package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Illuminateee/task-tracker.git/delivery/controller"
)

func main() {
	// Get the data file path (tasks.json in current directory)
	dataFilePath := getDataFilePath()

	// Create CLI controller
	cliController := controller.NewCLIController(dataFilePath)

	// Get command line arguments (skip program name)
	args := os.Args[1:]

	// Handle the command
	if err := cliController.HandleCommand(args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// getDataFilePath returns the path to the tasks.json file
func getDataFilePath() string {
	// Try to get from environment variable first
	if dataPath := os.Getenv("TASK_TRACKER_DATA"); dataPath != "" {
		return dataPath
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		// Fallback to current directory
		return "tasks.json"
	}

	// Return path to tasks.json in current directory
	return filepath.Join(cwd, "tasks.json")
}
