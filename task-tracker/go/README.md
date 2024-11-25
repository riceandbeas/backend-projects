<a name="readme-top"></a>

<div align="center">
<h3 align="center">Task Tracker</h3>
  <p align="center">
    CLI app to track and manage your tasks
    <br />
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About the project</a>
      <ul>
        <li><a href="#features">Features</a></li>
      </ul>
    </li>
    <li>
      <a href="#installation">Installation</a>
    </li>
    <li>
        <a href="#usage">Usage</a>
    </li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About the project

This is a solution for the [roadmap.sh](https://roadmap.sh/projects/task-tracker/solutions?u=673f5bde5434bf319a0b0302) Task Tracker project. Built in Go, it provides an intuitive command-line interface for managing tasks with elegantly styled output.

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Features

- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress
- Clean and visually appealing output

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- Installation -->
### Installation

Please note that this assumes you already have a Go environment set up. If not, you can check the documentation on the official [Go website](https://go.dev/doc/install).

1. Clone the repository and move into the project's directory:
  ```sh
  git clone https://github.com/riceandbeas/backend-projects.git
  cd backend-projects/task-tracker
  ```

2. Build the application:
  ```sh
  go build -o task-cli
  ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

Once the application is built, you can manage tasks through simple CLI commands. Below are some examples:

```sh
# Display help information about the available commands
./task-cli help

# Display help information about specified command
./task-cli help add

# Adding a new task
./task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
./task-cli update 1 "Buy groceries and cook dinner"
./task-cli delete 1

# Marking a task as in progress or done
./task-cli mark-in-progress 1
./task-cli mark-done 1

# Listing all tasks
./task-cli list

# Listing tasks by status
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
