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
        <a href="#usage">Usage</a>
    </li>
  </ol>
</details>



<!-- ABOUT THE PROJECT -->
## About the project

Task tracker is a [roadmap.sh project](https://roadmap.sh/projects/task-tracker/) from the [Backend Developer](https://roadmap.sh/backend) roadmap.

Solutions are available in the following languages:
- [Go](./go/)
- [Java](./java/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



### Features

- Add, Update, and Delete tasks
- Mark a task as in progress or done
- List all tasks
- List all tasks that are done
- List all tasks that are not done
- List all tasks that are in progress

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->
## Usage

```sh
# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>
