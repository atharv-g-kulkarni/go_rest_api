# go_rest_api

A REST API for managing events and user registrations, built with Go and the [Gin](https://github.com/gin-gonic/gin) web framework. Users can sign up, log in, create events, and register for events. Protected routes use JWT authentication.

## Live Deployment

The API is deployed on Render:

**https://go-rest-api-edhi.onrender.com/**

Use this as the base URL when calling endpoints (e.g. `GET https://go-rest-api-edhi.onrender.com/events`).

## Features

- User signup and login with bcrypt password hashing
- JWT-based authentication for protected endpoints
- CRUD operations for events (create, read, update, delete)
- Event registration and cancellation
- SQLite database with automatic schema migration on startup

## Tech Stack

- **Go** 1.25+
- **Gin** — HTTP router and middleware
- **SQLite** — local database (`api.db`)
- **JWT** — token-based auth (`golang-jwt/jwt`)
- **bcrypt** — password hashing

## Prerequisites

- Go 1.25 or later
- CGO enabled (required by `go-sqlite3`)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/atharv-g-kulkarni/go_rest_api.git
cd go_rest_api
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Configure environment variables

Create a `.env` file in the project root:

```env
secretKey=your-secret-key-here
```

This key is used to sign and verify JWT tokens.

### 4. Run the server

```bash
go run main.go
```

The server starts on `http://localhost:8080`. On first run, SQLite creates `api.db` and the required tables automatically.

## API Endpoints

Base URLs:

- **Production:** `https://go-rest-api-edhi.onrender.com`
- **Local:** `http://localhost:8080`

### Authentication

| Method | Path     | Auth | Description              |
|--------|----------|------|--------------------------|
| POST   | `/signup` | No  | Create a new user account |
| POST   | `/login`  | No  | Log in and receive a JWT  |

**Signup / Login request body:**

```json
{
  "EMAIL": "user@example.com",
  "PASSWORD": "yourpassword"
}
```

**Login response:**

```json
{
  "message": "Login successful!",
  "token": "<jwt-token>"
}
```

### Events

| Method | Path           | Auth | Description                    |
|--------|----------------|------|--------------------------------|
| GET    | `/events`      | No   | List all events                |
| GET    | `/events/:id`  | No   | Get a single event by ID       |
| POST   | `/events`      | Yes  | Create a new event             |
| PUT    | `/events/:id`  | Yes  | Update an event (owner only)   |
| DELETE | `/events/:id`  | Yes  | Delete an event (owner only)   |

**Create / Update event request body:**

```json
{
  "Name": "Go Meetup",
  "Description": "Monthly Go developers meetup",
  "Location": "San Francisco, CA",
  "DateTime": "2026-07-15T18:00:00Z"
}
```

### Registrations

| Method | Path                        | Auth | Description                         |
|--------|-----------------------------|------|-------------------------------------|
| POST   | `/events/:id/register`      | Yes  | Register the current user for an event |
| DELETE | `/events/:id/register`      | Yes  | Cancel registration for an event    |

### Authentication Header

Protected routes require a JWT in the `Authorization` header:

```
Authorization: <jwt-token>
```

## Project Structure

```
.
├── main.go              # Application entry point
├── db/
│   └── db.go            # SQLite connection and schema setup
├── models/
│   ├── users.go         # User model and database operations
│   └── events.go        # Event model and database operations
├── routes/
│   ├── routes.go        # Route registration
│   ├── users.go         # Signup and login handlers
│   ├── events.go        # Event CRUD handlers
│   └── register.go      # Event registration handlers
├── middlewares/
│   └── auth.go          # JWT authentication middleware
└── utils/
    ├── jwt.go           # Token generation and validation
    └── hash.go          # Password hashing utilities
```

## Database Schema

- **users** — `id`, `email` (unique), `password` (hashed)
- **events** — `id`, `name`, `description`, `location`, `dateTime`, `user_id`
- **registrations** — `id`, `event_id`, `user_id`
