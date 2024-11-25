import java.time.LocalDateTime;
import java.time.format.DateTimeFormatter;

public class Task {
    private final int id;
    private static int idCount = 0;
    String description;
    Status status;
    private String createdAt;
    private String updatedAt;

    public static final DateTimeFormatter dateTimeFormatter = DateTimeFormatter.ofPattern("dd-MM-yyyy HH:mm:ss");

    enum Status {
        TODO, INPROGRESS, DONE;

        public static Status fromString(String status) {
            return switch (status.toLowerCase().replace(" ", "")) {
                case "inprogress" -> INPROGRESS;
                case "done" -> DONE;
                default -> TODO;
            };
        }

        public String toString() {
            return switch (this) {
                case INPROGRESS -> "In progress";
                case DONE -> "Done";
                default -> "To do";
            };
        }
    }

    public Task(String description) {
        this.id = ++idCount;
        this.description = description;
        this.status = Status.TODO;
        this.createdAt = LocalDateTime.now().format(dateTimeFormatter);
        this.updatedAt = LocalDateTime.now().format(dateTimeFormatter);
    }

    public int getId() {
        return id;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public void setStatus(Status status) {
        this.status = status;
    }

    public String toString() {
        return """
                ------------------------------
                
                ID: %d
                Description: %s
                Status: %s
                Created at: %s
                Updated at: %s
                
                ------------------------------
                
                """.formatted(this.id, this.description, this.status, this.createdAt, this.updatedAt);
    }
}
