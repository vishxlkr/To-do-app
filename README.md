# Todo REST API

A RESTful Todo API built with **Go**, **Gin**, and **PostgreSQL**. Supports user authentication with JWT and full CRUD operations on todos — each user only sees and manages their own tasks.

---

## Tech Stack

| Layer      | Technology                   |
| ---------- | ---------------------------- |
| Language   | Go 1.25                      |
| Framework  | Gin                          |
| Database   | PostgreSQL                   |
| DB Driver  | pgx v5                       |
| Auth       | JWT (golang-jwt/jwt)         |
| Password   | bcrypt (golang.org/x/crypto) |
| Migrations | golang-migrate               |
| Config     | godotenv                     |

---

## Project Structure

```
.
├── cmd/
│   └── api/
│       └── main.go              # Entry point — wires config, DB, router
├── internal/
│   ├── config/
│   │   └── config.go            # Loads env variables
│   ├── database/
│   │   └── postgres.go          # PostgreSQL connection pool
│   ├── handlers/
│   │   ├── todo_handler.go      # Todo HTTP handlers
│   │   └── user_handler.go      # Auth HTTP handlers (register / login)
│   ├── middleware/
│   │   └── auth_middleware.go   # JWT validation middleware
│   ├── models/
│   │   ├── todo.go              # Todo database model
│   │   └── user.go              # User database model
│   └── repository/
│       ├── todo_repository.go   # Todo SQL queries
│       └── user_repository.go   # User SQL queries
├── migrations/                  # SQL migration files (up/down)
├── scripts/
│   └── migrate.ps1              # PowerShell helper for running migrations
├── go.mod
└── go.sum
```

---

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) 1.21+
- [PostgreSQL](https://www.postgresql.org/)
- [golang-migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

```powershell
# Install migrate CLI (PostgreSQL tag)
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 1. Clone & Install Dependencies

```bash
git clone <repo-url>
cd "project4 - Todo"
go mod tidy
```

### 2. Configure Environment

Copy the example env file and fill in your values:

```bash
cp .env.example .env
```

```env
DATABASE_URL=postgres://user:password@localhost:5432/todo_db?sslmode=disable
PORT=8080
JWT_SECRET=your_super_secret_key
```

### 3. Run Migrations

```powershell
# Apply all migrations
.\scripts\migrate.ps1 up

# Roll back the last migration
.\scripts\migrate.ps1 down

# Roll back N migrations
.\scripts\migrate.ps1 down 2

# Create a new migration
.\scripts\migrate.ps1 create migration_name
```

### 4. Start the Server

```bash
go run ./cmd/api
```

The API will be available at `http://localhost:8080`.

---

## API Reference

### Health Check

| Method | Endpoint | Auth | Description         |
| ------ | -------- | ---- | ------------------- |
| GET    | `/`      | No   | Server health check |

**Response**

```json
{
   "message": "Todo API is running",
   "success": true,
   "database": "connected"
}
```

---

### Auth

| Method | Endpoint         | Auth | Description          |
| ------ | ---------------- | ---- | -------------------- |
| POST   | `/auth/register` | No   | Register a new user  |
| POST   | `/auth/login`    | No   | Login, returns a JWT |

**Register — Request Body**

```json
{
   "email": "user@example.com",
   "password": "securepassword"
}
```

**Login — Request Body**

```json
{
   "email": "user@example.com",
   "password": "securepassword"
}
```

**Login — Response**

```json
{
   "token": "<jwt_token>"
}
```

---

### Todos (Protected)

All todo routes require the `Authorization` header:

```
Authorization: Bearer <jwt_token>
```

| Method | Endpoint     | Description                  |
| ------ | ------------ | ---------------------------- |
| POST   | `/todos`     | Create a new todo            |
| GET    | `/todos`     | Get all todos for the user   |
| GET    | `/todos/:id` | Get a specific todo by ID    |
| PUT    | `/todos/:id` | Update a todo (title/status) |
| DELETE | `/todos/:id` | Delete a todo                |

**Create Todo — Request Body**

```json
{
   "title": "Buy groceries",
   "completed": false
}
```

**Update Todo — Request Body** _(all fields optional)_

```json
{
   "title": "Buy groceries and cook",
   "completed": true
}
```

**Todo Object**

```json
{
   "id": 1,
   "title": "Buy groceries",
   "completed": false,
   "user_id": "uuid-of-owner",
   "created_at": "2026-03-15T10:00:00Z",
   "updated_at": "2026-03-15T10:00:00Z"
}
```

---

## Database Schema

### `users`

| Column     | Type                     | Notes            |
| ---------- | ------------------------ | ---------------- |
| id         | UUID                     | Primary Key      |
| email      | VARCHAR(255)             | Unique, Not Null |
| password   | VARCHAR(255)             | bcrypt hashed    |
| created_at | TIMESTAMP WITH TIME ZONE |                  |
| updated_at | TIMESTAMP WITH TIME ZONE |                  |

### `todos`

| Column     | Type         | Notes               |
| ---------- | ------------ | ------------------- |
| id         | SERIAL       | Primary Key         |
| title      | VARCHAR(255) | Not Null            |
| completed  | BOOLEAN      | Default: false      |
| user_id    | UUID         | Foreign key → users |
| created_at | TIMESTAMP    |                     |
| updated_at | TIMESTAMP    |                     |

---

## Architecture

```
Client
  │
  ▼
Gin Router
  │
  ├── Public Routes ─────────────► Auth Handlers ──► User Repository ──► PostgreSQL
  │
  └── Protected Routes ──► Auth Middleware (JWT) ──► Todo Handlers ──► Todo Repository ──► PostgreSQL
```

- **Handlers** parse HTTP requests and return JSON responses.
- **Repository** layer contains all raw SQL queries, keeping handlers clean.
- **Middleware** validates the JWT and injects the `user_id` into the request context, so handlers never deal with tokens directly.
- **Config** is loaded from environment variables via `.env`.
