# go-tasks-api

Go REST API paired with a Flutter client — run locally with list, create, and complete.

**Pair:** [flutter-task-app](https://github.com/wmsing/flutter-task-app) — Flutter client for this API.

Portfolio sample, not production.

## Run (≈2 min)

1. Clone and enter this repo.
2. Start the server:

```bash
go run ./cmd/server
```

Optional: copy `.env.example` to `.env` and set `PORT` (default `8080`).

3. Verify:

```bash
curl -s http://127.0.0.1:8080/health
curl -s http://127.0.0.1:8080/tasks
curl -s -X POST http://127.0.0.1:8080/tasks \
  -H 'Content-Type: application/json' \
  -d '{"title":"Demo task"}'
```

Replace `task-1` with the `id` from the create response:

```bash
curl -s -X PATCH http://127.0.0.1:8080/tasks/task-1/complete
```

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/health` | Liveness |
| GET | `/tasks` | List tasks |
| POST | `/tasks` | Create `{ "title": "..." }` |
| PATCH | `/tasks/:id/complete` | Mark complete (idempotent) |

Errors use JSON: `{"error":"..."}`.

## Quality

```bash
go fmt ./...
go vet ./...
go test ./...
```

## Stack

Go 1.22+ (`net/http` routing), in-memory store.
