# Task Tracker
The **task tracker** is a simple project created following the instruction from the [task-tracker](https://roadmap.sh/projects/task-tracker)  challenge, available on the [Backend Developer](https://roadmap.sh/backend) roadmap.<br>

## How to use it
### Installing:
Clone this repository as the command bellow says:
```
git clone https://github.com/vianavitor-dev/task-tracker.git
```
### Usage:
There are the available commands that this project has:
``` golang
# use this to go on the project directory
cd task-tracker

# Adding a new Task
go run main.go add "Buy groceries"

# Updating a Task
go run main.go update 1 "Buy groceries and cook dinner"

# Deleting a Task
go run main.go delete 1

# Marking a Task as in progress or done
go run main.go mark-in-progress 1
go run main.go mark-done 1

# Listing all Tasks
go run main.go list

# Listing Tasks by status
go run main.go list done
go run main.go list todo
go run main.go list in-progress
```
<br>
