# Tasker CLI

A project and taks management kanban board in your CLI

Many developers thrive in the focused environment of the command line. Switching to browsers or other apps can disrupt their workflow and lead to distractions.
Tasker is a simple CLI TUI (terminal user interface) application designed specifically for developers. It helps you stay focused on your current projects and their associated tasks by eliminating the distractions and overhead of managing separate apps.
Inspired by Vim keybindings, Tasker offers an intuitive user experience. Additionally, BubbleTea allows it to maintain a visually appealing interface.
If you find yourself getting sidetracked while working on personal projects, Tasker can be a valuable addition to your development toolkit.


## Background
Tasker is actually a proof of concept developed over a weekend, inspired by the BubbleTea guide on creating TUI apps.
It originated when I was looking for side projects to work on and develop new skills. I realized that I easily get distracted, so I decided to build something that would hold me accountable whenever I start a new project.
After learning about BubbleTea and the convenience of terminal applications, I decided to start my learning journey with Tasker CLI.

## Install
This project use [Go 1.22](https://go.dev/dl/)
```bash
git clone https://github.com/Desgue/Tasker-CLI.git
```
Then run on the cloned dir
```bash
go mod tidy
```
Build and Install
```bash
go build && go install
```
## Running
Then simply run 
```bash
tasker-cli
```
If you dont wish to install Into your path just run it from root
```bash
go run .
```

## Views

**Project View:**

The project view displays your projects organized as Kanban boards. 
Each board represents a priority status and contains projects categorized by their priority (e.g., Low, Medium, High). 
Selecting a project board allows you to view and interact with its associated tasks, such as opening the task kanban board and creating new tasks.

**Task View:**

The task view displays your tasks organized as Kanban boards. 
Each board represents a status and contains tasks categorized by their status (e.g., Pending, In Progress, Done). 

## Commands

### Project View Keybindings

| Keybinding | Action |
|---|---|
| `q`, `ctrl+c` | Quit the program |
| `l`, `right`, `tab` | Switch to the next priority board |
| `h`, `left` | Switch to the previous priority board |
| `d`, `delete` | Delete the selected project |
| `backspace`, `ctrl+b` | Decrease the selected task's priority (if any) |
| `space`, `ctrl+n` | Increase the selected task's priority (if any) |
| `n` | Open the form to create a new project |
| `t` | Open the tasks for the selected project (if any) |




### Task View Keybindings

| Keybinding | Action |
|---|---|
| `q`, `ctrl+c` | Quit the program |
| `esc` | Go back to the project view |
| `l`, `right`, `tab` | Switch to the next status board |
| `h`, `left` | Switch to the previous status board|
| `d`, `delete` | Delete the selected task |
| `backspace`, `ctrl+b` | Decrease the selected task's status (if any) |
| `space`, `ctrl+n` | Increase the selected task's status (if any) |
| `n` | Open the form to create a new task |

**Notes:**

* Some actions require a project or task to be selected in the currently focused board.
* `ctrl+b` and `ctrl+n` are keyboard shortcuts achieved by holding the `ctrl` key while pressing `b` or `n`, respectively.

