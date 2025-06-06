# go-api-cms

This project is a backend API built with Golang, following a clean architecture approach. The project structure is organized into several layers to ensure modularity and scalability.

## 🚀 How to Run the Application

### 📌 Notes
Make sure you have a valid .env file (based on .env.example) before running the application.

### Run in Development Mode
```bash
go run .
```

### Run Database Migrations
```bash
go run . db:migrate
```

### Run Database Seeders
```bash
go run . db:migrate
```

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
