# POS System - Clean Architecture

Point of Sale (POS) system built with Go using Clean Architecture principles.

## Tech Stack

- **Framework**: Go Fiber
- **ORM**: GORM
- **Database**: PostgreSQL
- **Cache**: Redis
- **Storage**: MinIO
- **Migration**: golang-migrate
- **Logger**: Zap
- **Validator**: go-playground/validator
- **JWT**: golang-jwt

## Project Structure
```
pos-system/
├── cmd/                    # Application entry points
├── config/                 # Configuration management
├── internal/               # Private application code
│   ├── domain/            # Business logic & entities
│   ├── usecase/           # Application business rules
│   ├── delivery/          # Handlers, middleware, routes
│   ├── repository/        # Data access implementation
│   └── infrastructure/    # External services setup
├── pkg/                   # Public libraries
├── database/              # Migrations
└── seeders/              # Database seeders
```

## Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 15+
- Redis 7+
- MinIO
- golang-migrate CLI

### Installation

1. Clone the repository
```bash
git clone <repository-url>
cd pos-system
```

2. Setup project
```bash
make setup
```

3. Configure environment variables
```bash
# Edit .env file with your configuration
nano .env
```

4. Start dependencies with Docker
```bash
docker-compose up -d
```

5. Run migrations
```bash
make migrate-up
```

6. Seed database
```bash
make seed
```

7. Run application
```bash
# Development mode (hot reload)
make dev

# Production mode
make run
```

## Available Commands

### Development
```bash
make dev              # Run with hot reload
make run              # Run application
make build            # Build binary
```

### Database
```bash
make migrate-up       # Run migrations
make migrate-down     # Rollback migrations
make migrate-create   # Create new migration
make migrate-force    # Force migration version
make seed             # Run seeders
make db-reset         # Reset database (down, up, seed)
```

### Testing
```bash
make test             # Run tests
make test-coverage    # Run tests with coverage
```

### Utilities
```bash
make clean            # Clean temporary files
make install          # Install dependencies
```

## Environment Variables

See `.env.example` for all available configuration options.

## API Documentation

API documentation will be available at:
- Swagger UI: `http://localhost:8080/swagger`

## Default Credentials

After running seeders:
- **Admin**: admin@pos.com / password123
- **Cashier**: cashier1@pos.com / password123
- **Manager**: manager@pos.com / password123

## License

MIT License
```

---

## **Struktur Folder Final:**
```
pos-system/
│
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── seeder/
│       └── main.go
│
├── config/
│   ├── config.go
│   └── loader.go
│
├── internal/
│   ├── domain/
│   │   ├── entity/
│   │   └── repository/
│   ├── usecase/
│   ├── delivery/
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   ├── middleware/
│   │   │   └── routes/
│   │   └── dto/
│   │       ├── request/
│   │       └── response/
│   ├── repository/
│   │   ├── gorm/
│   │   └── redis/
│   └── infrastructure/
│       ├── database/
│       ├── cache/
│       ├── storage/
│       └── fiber/
│
├── pkg/
│   ├── logger/
│   ├── jwt/
│   ├── validator/
│   ├── response/
│   ├── pagination/
│   ├── utils/
│   └── constants/
│
├── database/
│   └── migrations/
│
├── seeders/
│
├── docs/
│
├── tests/
│
├── scripts/
│
├── storage/
│   ├── logs/
│   └── temp/
│
├── .env
├── .env.example
├── .gitignore
├── .air.toml
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
└── README.md