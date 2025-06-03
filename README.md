# notesAPI

A simple RESTful API built in Go to manage notes.
This is my first CRUD API written in [Go](https://go.dev/).

> [!NOTE]
> This project is not completed yet.

## Documentation

### Features

- A simple RESTful API with CRUD operations.
- JSON-based request and response.
- SQLite database.

### API Endpoints

| Method | Endpoint     | Description       |
| ------ | ------------ | ----------------- |
| GET    | `/notes`     | Get all notes     |
| GET    | `/notes/:id` | Get note by ID    |
| POST   | `/notes`     | Create a new note |
| PUT    | `/notes/:id` | Update a note     |
| DELETE | `/notes/:id` | Delete a note     |

### Request Body

```json
{
  "id": 1,
  "title": "Meeting Notes",
  "content": "Discuss project timelines.",
  "created_at": "2025-06-03T15:00:00Z",
  "updated_at": "2025-06-03T15:00:00Z"
}
```
