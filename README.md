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

# Struktur Folder Clean Architecture POS System

```
pos-system/
│
├── cmd/
│   └── api/
│       └── main.go                          # Entry point aplikasi
│
├── config/
│   ├── config.go                            # Struct config
│   └── loader.go                            # Viper .env loader
│
├── internal/
│   │
│   ├── domain/                              # Enterprise Business Rules
│   │   ├── entity/
│   │   │   ├── user.go
│   │   │   ├── product.go
│   │   │   ├── category.go
│   │   │   ├── transaction.go
│   │   │   ├── transaction_item.go
│   │   │   ├── customer.go
│   │   │   ├── payment.go
│   │   │   └── base.go                      # GORM base model
│   │   │
│   │   └── repository/                      # Interface repository
│   │       ├── user_repository.go
│   │       ├── product_repository.go
│   │       ├── category_repository.go
│   │       ├── transaction_repository.go
│   │       ├── customer_repository.go
│   │       └── cache_repository.go          # Redis interface
│   │
│   ├── usecase/                             # Application Business Rules
│   │   ├── product/
│   │   │   ├── create_product.go
│   │   │   ├── update_product.go
│   │   │   ├── delete_product.go
│   │   │   ├── get_product.go
│   │   │   └── upload_image.go              # MinIO integration
│   │   │
│   │   ├── transaction/
│   │   │   ├── create_transaction.go
│   │   │   ├── get_transaction.go
│   │   │   ├── get_transaction_list.go
│   │   │   └── generate_report.go
│   │   │
│   │   ├── customer/
│   │   │   ├── create_customer.go
│   │   │   ├── update_customer.go
│   │   │   └── get_customer.go
│   │   │
│   │   ├── category/
│   │   │   ├── create_category.go
│   │   │   └── get_category.go
│   │   │
│   │   └── auth/
│   │       ├── login.go
│   │       ├── register.go
│   │       ├── refresh_token.go
│   │       └── logout.go
│   │
│   ├── delivery/                            # Interface Adapters
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   │   ├── product_handler.go
│   │   │   │   ├── transaction_handler.go
│   │   │   │   ├── customer_handler.go
│   │   │   │   ├── category_handler.go
│   │   │   │   ├── auth_handler.go
│   │   │   │   └── upload_handler.go        # File upload handler
│   │   │   │
│   │   │   ├── middleware/
│   │   │   │   ├── auth.go                  # JWT middleware
│   │   │   │   ├── cors.go
│   │   │   │   ├── logger.go                # Zap logger middleware
│   │   │   │   ├── rate_limiter.go          # Redis rate limiter
│   │   │   │   ├── error_handler.go
│   │   │   │   └── request_id.go
│   │   │   │
│   │   │   └── routes/
│   │   │       ├── routes.go                # Main router
│   │   │       ├── api.go                   # API routes
│   │   │       └── public.go                # Public routes
│   │   │
│   │   └── dto/                             # Data Transfer Objects
│   │       ├── request/
│   │       │   ├── auth_request.go
│   │       │   ├── product_request.go
│   │       │   ├── transaction_request.go
│   │       │   ├── customer_request.go
│   │       │   ├── category_request.go
│   │       │   └── pagination_request.go
│   │       │
│   │       └── response/
│   │           ├── auth_response.go
│   │           ├── product_response.go
│   │           ├── transaction_response.go
│   │           ├── customer_response.go
│   │           ├── category_response.go
│   │           ├── pagination_response.go
│   │           └── base_response.go
│   │
│   ├── repository/                          # Frameworks & Drivers
│   │   ├── gorm/
│   │   │   ├── user_repository.go
│   │   │   ├── product_repository.go
│   │   │   ├── category_repository.go
│   │   │   ├── transaction_repository.go
│   │   │   └── customer_repository.go
│   │   │
│   │   └── redis/
│   │       ├── cache_repository.go
│   │       └── session_repository.go
│   │
│   └── infrastructure/                      # External Services Setup
│       ├── database/
│       │   ├── gorm.go                      # GORM setup & connection
│       │   └── migrate.go                   # golang-migrate setup
│       │
│       ├── cache/
│       │   └── redis.go                     # Redis setup & connection
│       │
│       ├── storage/
│       │   └── minio.go                     # MinIO setup & connection
│       │
│       └── fiber/
│           └── fiber.go                     # Fiber app configuration
│
├── pkg/                                      # Shared Utilities
│   ├── logger/
│   │   ├── logger.go                        # Zap logger wrapper
│   │   └── logger_test.go
│   │
│   ├── jwt/
│   │   ├── jwt.go                           # JWT service
│   │   └── claims.go                        # JWT claims
│   │
│   ├── validator/
│   │   ├── validator.go                     # Validator v10 wrapper
│   │   └── custom_validator.go             # Custom validations
│   │
│   ├── response/
│   │   ├── response.go                      # Standard response helper
│   │   └── error.go                         # Error response helper
│   │
│   ├── pagination/
│   │   └── pagination.go                    # Pagination helper
│   │
│   ├── utils/
│   │   ├── hash.go                          # Password hashing
│   │   ├── string.go                        # String utilities
│   │   ├── time.go                          # Time utilities
│   │   └── file.go                          # File utilities
│   │
│   └── constants/
│       ├── error.go                         # Error constants
│       ├── message.go                       # Message constants
│       └── role.go                          # Role constants
│
├── database/
│   └── migrations/                          # golang-migrate migrations
│       ├── 000001_create_users_table.up.sql
│       ├── 000001_create_users_table.down.sql
│       ├── 000002_create_categories_table.up.sql
│       ├── 000002_create_categories_table.down.sql
│       ├── 000003_create_products_table.up.sql
│       ├── 000003_create_products_table.down.sql
│       ├── 000004_create_customers_table.up.sql
│       ├── 000004_create_customers_table.down.sql
│       ├── 000005_create_transactions_table.up.sql
│       ├── 000005_create_transactions_table.down.sql
│       ├── 000006_create_transaction_items_table.up.sql
│       ├── 000006_create_transaction_items_table.down.sql
│       ├── 000007_create_payments_table.up.sql
│       └── 000007_create_payments_table.down.sql
│
├── seeders/                                 # Database seeders
│   ├── seeder.go                            # Main seeder runner
│   ├── user_seeder.go
│   ├── category_seeder.go
│   └── product_seeder.go
│
├── docs/
│   ├── swagger/                             # Swagger documentation
│   │   └── swagger.json
│   │
│   └── api/                                 # API documentation
│       └── README.md
│
├── tests/
│   ├── unit/
│   │   ├── usecase/
│   │   │   └── product_test.go
│   │   │
│   │   └── repository/
│   │       └── product_repository_test.go
│   │
│   ├── integration/
│   │   └── api/
│   │       └── product_api_test.go
│   │
│   └── mocks/
│       ├── product_repository_mock.go
│       └── cache_repository_mock.go
│
├── scripts/
│   ├── setup.sh                             # Setup script
│   ├── migrate-up.sh                        # Run migrations up
│   ├── migrate-down.sh                      # Run migrations down
│   ├── migrate-create.sh                    # Create new migration
│   └── seed.sh                              # Seeder script
│
├── storage/                                 # Local storage (gitignored)
│   ├── logs/
│   └── temp/
│
├── .env.example
├── .env
├── .gitignore
├── .air.toml                                # Air hot reload config
├── docker-compose.yml                       # Docker compose setup
├── Dockerfile
├── Makefile
├── go.mod
├── go.sum
└── README.md
```