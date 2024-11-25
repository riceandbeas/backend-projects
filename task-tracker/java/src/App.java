public class App {
    Tracker tracker = new Tracker();

    void addTask(String description) {
        int id = this.tracker.addTask(description);
        System.out.println("Task added successfully (ID: " + id + ").");
    }

    void listTasks(String filter) {
        this.tracker.listTasks(filter);
    }

    void updateTask(int id, String description) {
        tracker.updateTask(id, description);
    }

    private void printHelp() {
        System.out.println("""
                Available commands:
                - help: Displays the list of available commands.

                - add: Adds a new task. Requires one additional argument:
                    - description

                - list: Lists tasks. Optionally takes one additional argument:
                    - status (optional)

                - update: Renames an existing task. Requires two additional arguments:
                    - id
                    - description

                - delete: Deletes a task. Requires one additional argument:
                    - id
                """);
    }

    public static void main(String[] args) {
        if (args.length == 0) {
            System.out.println("Please provide a command. Use 'help' for a list of commands.");
            return;
        }

        App app = new App();

        String cmd = args[0];
        switch (cmd.toLowerCase()) {
            case "help":
                app.printHelp();
                break;

            case "add":
                if (args.length < 2 ) {
                    System.out.println("Please provide a description for the task.");
                    return;
                }

                app.addTask(args[1]);
                break;

            case "update":
                if (args.length < 3) {
                    System.out.println("Please provide an ID and the new description");
                    return;
                }

                app.updateTask(Integer.parseInt(args[1]), args[2]);
                break;

            case "list":
                if (args.length < 2) {
                    app.listTasks("");
                } else {
                    app.listTasks(args[1]);
                }

                break;

            default:
                System.out.println("Invalid command. Use 'help' for a list of commands.");
        }
        app.tracker.saveTasks();
    }
}
