# go-eCommerce-core

A domain-driven design (DDD) backend service built in Go using the standard net/http package. The project implements core CRUD operations for products and users and follows a clean, layered architecture (domain, repository, infra, REST handlers). This README was prepared from an inspection of the repository layout, module information, and the configuration example included in the project.

---

## Table of contents

- About
- Project structure
- Key files inspected
- Requirements
- Configuration (.env)
- Build & Run
- Database & Migrations
- API (overview)
- Authentication
- Diagnostics & logs
- Contributing
- License

---

## About

go-eCommerce-core is a small, opinionated backend service implementing core concepts of CURD, DDD, JWT in Go. It uses:

- Go 1.25 (module `ecommerece`)
- net/http for the HTTP server
- sqlx + lib/pq for PostgreSQL access
- godotenv for environment configuration

The service aims to provide CRUD operations for products and users with separation of concerns across directories such as `domain`, `repo`, `rest`, `product`, `user`, `infra`, and `config`.

---

## Project structure (high level)

These are the top-level folders in the repository:

- cmd — application entrypoint(s). The server is started via `cmd.Serve()` (called from `main.go`).
- config — configuration loading and structs.
- infra — infrastructure code (database connection, logging, etc).
- domain — domain models and business logic interfaces.
- repo — repository implementations (data access).
- rest — HTTP handlers / routing.
- product — product-specific application layer (service, handlers, models).
- user — user-specific application layer (service, handlers, models).
- db_queries — SQL queries used by repository layer.
- migrations — database migration files / SQL.
- util — utility helpers.
- main.go — simple entrypoint that calls `cmd.Serve()`.

Note: The above is based on the repository tree; consult the code in each directory for concrete types and signatures.

---

## Key files inspected

- main.go
  - Entrypoint that calls `cmd.Serve()`.
- go.mod
  - Module: `ecommerece`
  - Declares dependencies: github.com/jmoiron/sqlx, github.com/joho/godotenv, github.com/lib/pq, etc.
- .env.example
  - Contains the environment variables expected by the application (see next section).

---

## Requirements

- Go 1.25.x (matches go.mod)
- PostgreSQL

---

## Configuration

Copy `.env.example` to `.env` and fill in values:

Example `.env` entries (from .env.example):

VERSION=1.0.0  
SERVICE_NAME=ECOMMERCE  
HTTP_PORT=3000  
JWT_SECRET_KEY=change-<this-to-a-very-long-random-string-at-least-32-chars>

DB_HOST=localhost  
DB_PORT=5432  
DB_NAME=ecommerce  
DB_USER=<user_name>
DB_PASSWORD=<your_strong_password_here>  
DB_ENABLE_SSL_MODE=false

Notes:
- Replace with the correct values in your environment.

---

## Build & Run

From the repository root:

1. Install dependencies (modules are used; this happens automatically when you build)
   - go mod download

2. Run directly (for development):
   - go run main.go

   or build and run:

   - go build -o bin/ecommerce ./  
   - ./bin/ecommerce

By default the service listens on HTTP_PORT from the environment (default: 3000).

The `main.go` entrypoint calls `cmd.Serve()` which wires up configuration, database connections, and HTTP handlers.

---

## Database & Migrations

This project includes a `migrations` folder and `db_queries` folder. Inspect the migration SQL files and run them against your Postgres instance before starting the service.

Example (psql):

- Create the database:
  - createdb -h $DB_HOST -p $DB_PORT -U $DB_USER $DB_NAME
- Run migration files:
  - psql "postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable" -f migrations/0001_init.sql
  - Repeat for other migration files in the `migrations/` folder.

If the repository includes a migration tool wrapper, use it (for instance, a `migrate` binary or an included script). If not, apply SQL files in order manually.

---

## API (overview)

The codebase implements REST endpoints for products and users following typical CRUD operations. Exact routes and JSON schemas are defined in the `rest`, `product`, and `user` packages; inspect those handlers for precise routes and request/response shapes.

Typical operations you can expect:

Products
- GET /products — list products (with optional pagination/filters)
- GET /products/{id} — get a single product by id
- POST /products — create a new product
- PUT /products/{id} — update a product
- DELETE /products/{id} — delete a product

Users
- POST /users/register — create/register new user
- POST /users/login — authenticate and receive JWT
- GET /users/{id} — get user details
- PUT /users/{id} — update user
- DELETE /users/{id} — delete user

After login you should receive a JWT token which you attach in Authorization header as:
  Authorization: Bearer <token>

---

## Authentication

- The project expects a JWT secret (JWT_SECRET_KEY) in environment.
- Authentication is likely implemented for user login and protected endpoints. Check the `rest` and `user` packages for middleware/handlers that enforce JWT authentication.

---

## Diagnostics & logs

- Check logs written to stdout/stderr by the server; the `infra` package likely contains logger setup.
- Database errors and SQL queries are routed through the repository layer (sqlx).

---

