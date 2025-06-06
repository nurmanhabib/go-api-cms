# go-api-cms

This project is a backend API built with Golang, following a clean architecture approach. The project structure is organized into several layers to ensure modularity and scalability.

## 🚀 Running the Application with Docker Compose

This application can be run entirely using **Docker Compose**, keeping your host environment clean and untouched. All dependencies (such as the database) will run inside isolated containers.

However, to avoid port conflicts with existing services on your host machine, take note of the following:

📌 Important Environment Variables
- `FORWARD_DB_PORT`: The port on your host that forwards to the PostgreSQL service inside the container.
  - Default: `5432`
  - If you already have PostgreSQL running on port 5432 on your host, change this to another port like 5433.
- `APP_PORT`: The port on your host used to access the application.
  - Example: `8080`

Create a .env file in the project root:
```
FORWARD_DB_PORT=5433
APP_PORT=8080
```

### ⚡️ Minimal Configuration – Just Works

You can run the application with a single command:

```bash
docker compose up -d
```

### 🛑 Stopping the Application

To stop and remove containers:
```bash
docker compose down
```

## Run in Development Mode
```bash
go run .
```

### Run Database Migrations
```bash
go run . db:migrate
```

### Run Database Seeders
```bash
go run . db:seed
```

## 🧪 Database Seeding

Once the application starts and the database seeder runs, the following pre-defined roles will be available:
- admin
- editor

These roles are created automatically during the initial database seeding process.
You can assign these roles to users for role-based access control (RBAC) in the application.

## 📁 Project Directory Structure
```
go-api-cms/
├── app/                      # Application initialization and dependency injection
├── config/                   # Application configuration files
├── database/                 # Database configuration
│   ├── migrations/           # Migration files
│   └── seeder/               # Seeder files
├── domain/                   # Domain layer (business logic)
│   ├── entity/               # Entity definitions
│   └── repository/           # Repository interfaces
├── infrastructure/           # Implementations of domain interfaces
│   ├── dao/                  # Data Access Objects (DAO)
│   ├── exception/            # Error and exception handling
│   ├── persistence/          # Data persistence implementations
│   └── services/             # External or internal services
├── interfaces/               # Interfaces to the outside world (API, CLI, etc.)
│   ├── cmd/                  # Command Line Interface (CLI)
│   └── http/                 # HTTP handling
│       ├── controller/       # HTTP controllers
│       └── middleware/       # HTTP middleware
├── internal/                 # Internal utilities
│   ├── basepath/             # Base path utility
│   └── version/              # Application version info
├── pkg/                      # Reusable libraries or utilities
│   ├── env/                  # Environment variable management
│   ├── graceful/             # Graceful shutdown handling
│   └── provider/             # External providers
├── routes/                   # Route definitions
│   └── api.go
├── .editorconfig             # Editor configuration
├── .env                      # Environment variable file
├── .env.example              # Example environment file
├── .gitignore                # Git ignored files
├── go.mod                    # Go module file
└── main.go                   # Application entry point
```

## 🧱 Technologies
- Golang
- Clean Architecture
- GORM (ORM)
- HTTP + Middleware
- CLI for migration & seeding
