# Task Tracker CLI

A simple and efficient command-line task management application written in Go. This tool helps you organize, track, and manage your tasks with full CRUD operations and status filtering capabilities.

## Features

- ✅ **Full CRUD Operations**: Create, Read, Update, and Delete tasks
- 📁 **JSON Storage**: Tasks are stored in a JSON file for persistence
- 🏷️ **Status Management**: Track tasks with three statuses: `todo`, `in-progress`, `done`
- 🔍 **Smart Filtering**: List tasks by status or view all tasks
- 📋 **Pending Tasks View**: See all non-completed tasks at once
- 🕒 **Timestamps**: Automatic creation and update timestamps
- 💻 **Cross-platform**: Works on Windows, macOS, and Linux

## Installation

### Prerequisites
- Go 1.25+ installed on your system

### Build from Source
1. Clone the repository:
```bash
git clone https://github.com/Illuminateee/task-tracker.git
cd task-tracker
```

2. Build the application:
```bash
go build -o task-tracker cmd/main.go
```

### Windows Users
```powershell
go build -o task-tracker.exe cmd/main.go
```

## Usage

### Basic Commands

#### Add a New Task
```bash
# Add task with title only
./task-tracker add "Buy groceries"

# Add task with title and description
./task-tracker add "Buy groceries" "Milk, bread, eggs, cheese"
```

#### List Tasks
```bash
# List all tasks
./task-tracker list
./task-tracker list all

# List completed tasks
./task-tracker list done

# List todo tasks
./task-tracker list todo

# List in-progress tasks
./task-tracker list in-progress

# List all pending tasks (todo + in-progress)
./task-tracker list pending
```

#### Update Tasks
```bash
# Update task title and description
./task-tracker update 1 "Buy organic groceries" "Organic milk, whole grain bread, free-range eggs"

# Update only title (description remains unchanged)
./task-tracker update 1 "Buy organic groceries"
```

#### Change Task Status
```bash
# Mark task as done
./task-tracker mark-done 1

# Mark task as in progress
./task-tracker mark-in-progress 1

# Mark task as todo
./task-tracker mark-todo 1
```

#### Delete Tasks
```bash
# Delete a task by ID
./task-tracker delete 1
```

#### Get Help
```bash
./task-tracker help
```

## Task Statuses

| Status | Description |
|--------|-------------|
| `todo` | Task is created but not started |
| `in-progress` | Task is currently being worked on |
| `done` | Task is completed |

## Data Storage

Tasks are stored in a JSON file called `tasks.json` in the current working directory. You can customize the storage location by setting the `TASK_TRACKER_DATA` environment variable:

```bash
# Linux/macOS
export TASK_TRACKER_DATA="/path/to/your/tasks.json"

# Windows
set TASK_TRACKER_DATA="C:\path\to\your\tasks.json"
```

### Sample JSON Structure
```json
[
  {
    "id": 1,
    "title": "Buy groceries",
    "description": "Milk, bread, eggs, cheese",
    "status": "done",
    "created_at": "2025-10-06T10:30:00Z",
    "updated_at": "2025-10-06T15:45:00Z"
  },
  {
    "id": 2,
    "title": "Write project documentation",
    "description": "Create comprehensive README and API docs",
    "status": "in-progress",
    "created_at": "2025-10-06T11:00:00Z",
    "updated_at": "2025-10-06T14:20:00Z"
  }
]
```

## Examples

### Complete Workflow Example
```bash
# Add some tasks
./task-tracker add "Complete project proposal" "Draft, review, and submit"
./task-tracker add "Schedule team meeting" "Next week availability check"
./task-tracker add "Code review" "Review PR #123"

# Check all tasks
./task-tracker list

# Start working on a task
./task-tracker mark-in-progress 1

# Complete a task
./task-tracker mark-done 3

# Check pending work
./task-tracker list pending

# Update task details
./task-tracker update 2 "Schedule team meeting" "Check availability for next Tuesday"

# Remove completed task
./task-tracker delete 3
```

### Filtering Examples
```bash
# See only what needs to be done
./task-tracker list pending

# Check completed work
./task-tracker list done

# See current work in progress
./task-tracker list in-progress

# View all tasks for overview
./task-tracker list all
```

## Architecture

The application follows clean architecture principles:

```
cmd/
  main.go                    # Application entry point
delivery/
  controller/
    cli_controller.go        # CLI interface and command handling
entity/
  task.go                    # Task entity and business rules
manager/
  task_manager.go            # Application coordinator
repository/
  task_repository.go         # Repository interface
  json_task_repository.go    # JSON file implementation
usecase/
  task_usecase.go            # Business logic layer
```

## Error Handling

The application provides clear error messages for common scenarios:
- Invalid task IDs
- Missing required parameters
- File access issues
- Invalid status values
- Non-existent tasks

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature-name`
3. Make your changes
4. Add tests if applicable
5. Commit your changes: `git commit -am 'Add feature'`
6. Push to the branch: `git push origin feature-name`
7. Submit a pull request

## Development

### Running Tests
```bash
go test ./...
```

### Code Structure
- **Entity Layer**: Core business entities and rules
- **Repository Layer**: Data persistence abstraction
- **Use Case Layer**: Business logic implementation
- **Manager Layer**: Application coordination
- **Delivery Layer**: User interface (CLI)

## License

This project is open source and available under the [MIT License](LICENSE).

## Changelog

### v1.0.0
- Initial release
- Full CRUD operations
- JSON file storage
- Status filtering
- CLI interface

## Support

For issues, questions, or contributions, please visit the [GitHub repository](https://github.com/Illuminateee/task-tracker).