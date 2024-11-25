import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.io.Reader;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;
import java.util.stream.Stream;

public class Tracker {
    List<Task> tasks = new ArrayList<>();
    private final Path TASKS_PATH = Path.of("tasks.json");
    private final Gson gson = new Gson();

    Tracker() {
        this.tasks = loadTasks();
    }

    private List<Task> loadTasks() {
        List<Task> tasks = new java.util.ArrayList<>();

        if (!Files.exists(TASKS_PATH)) {
            return tasks;
        }

        try {
            Reader reader = Files.newBufferedReader(TASKS_PATH);
            tasks = gson.fromJson(reader, new TypeToken<List<Task>>() {}.getType());
            reader.close();
            System.out.println(tasks);
        } catch (java.io.IOException e) {
            e.printStackTrace();
        }

        return tasks;
    }

    void saveTasks() {
        String json = gson.toJson(this.tasks);

        try {
            Files.writeString(TASKS_PATH, json);
        } catch (java.io.IOException e) {
            e.printStackTrace();
        }
    }

    int addTask(String description) {
        Task task = new Task(description);
        this.tasks.add(task);
        return task.getId();
    }

    void updateTask(int id, String description) {
        for (Task task : this.tasks) {
            if (task.getId() == id) {
                task.setDescription(description);
            }
        }
    }

    void listTasks(String filter) {
        if (!filter.isBlank()) {
            try {
                Stream<Task> stream = this.tasks.stream()
                        .filter(task -> task.status
                                .equals(Task.Status.fromString(filter)));
                stream.forEach(System.out::println);
            } catch (java.lang.NullPointerException e) {
                System.out.println("No tasks found.");
            }
        } else {
            System.out.println(this.tasks);
        }
    }
}
