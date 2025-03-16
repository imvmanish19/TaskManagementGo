# Task Service API

This is a simple task management service built using Go and the Gin framework. It supports operations like creating, updating, deleting, and retrieving tasks. Each task is associated with a user, and the task service communicates with a user service to retrieve user information.

## Problem Breakdown and Design Decisions

### Problem Breakdown

The task service is designed to handle tasks in a task management system. A task has the following properties:

- **ID**: Unique identifier for the task.
- **Name**: A brief title or name of the task.
- **Status**: The status of the task (e.g., "Pending", "In Progress", "Completed").
- **UserID**: The identifier of the user associated with the task.

The service provides the following functionalities:

1. **Create a Task**: Allows creating a new task.
2. **Get Tasks**: Retrieve a list of tasks, with pagination and optional filtering by status.
3. **Update a Task**: Update the details of an existing task.
4. **Delete a Task**: Delete a task by its ID.

### Design Decisions

- **In-Memory Storage**: Tasks are stored in-memory for simplicity. In a production environment, you should use a database to persist tasks.
- **Microservices Interaction**: The task service interacts with a **User Service** to fetch user information. This interaction is abstracted by an HTTP client that makes requests to the user service.
- **Pagination and Filtering**: The task service supports pagination to fetch a limited number of tasks per request and allows filtering tasks by status. These features are crucial for scalability.
- **Error Handling**: The task service gracefully handles errors, such as when the user service is unavailable or when a task is not found.

## Instructions to Run the Service

### Start Docker Engine

```bash
docker-compose up --build


