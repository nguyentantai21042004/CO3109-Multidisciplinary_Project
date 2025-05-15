# CO3109 Multidisciplinary Project - GoLang API

A comprehensive e-commerce platform backend API built with Go, featuring robust user management, shop management, and various third-party integrations.

## 🌟 Features

- **Authentication & Authorization**
  - JWT-based authentication
  - Role-based access control (RBAC)
  - Secure password handling

- **Core Functionalities**
  - User management system
  - Shop management
  - Province and element management
  - Shop user management

- **Integrations**
  - Email notifications via SMTP
  - Telegram bot integration
  - Cloudinary for media storage
  - Redis for caching
  - RabbitMQ for message queuing
  - Google OAuth2 integration

- **Development Tools**
  - Swagger API documentation
  - Docker containerization
  - Automated database migrations
  - Generated database models with SQLBoiler

## 🔧 Prerequisites

- Go 1.23.8 or higher
- Docker and Docker Compose
- PostgreSQL
- Make
- Redis
- RabbitMQ

## 📁 Project Structure

```
.
├── cmd/                # Application entry points
│   ├── api/           # Main API server
│   └── consumer/      # Message queue consumer
├── config/            # Configuration files
├── database/          # Database migrations and schema
├── docs/              # Swagger documentation
├── internal/          # Internal application code
│   ├── appconfig/     # Application configuration
│   ├── auth/          # Authentication module
│   ├── core/          # Core business logic
│   ├── element/       # Element management
│   ├── httpserver/    # HTTP server implementation
│   ├── middleware/    # HTTP middleware
│   ├── models/        # Generated database models
│   ├── province/      # Province management
│   ├── role/          # Role management
│   ├── shop/          # Shop management
│   ├── shopuser/      # Shop user management
│   └── user/          # User management
├── pkg/               # Reusable packages
└── vendor/            # Dependencies
```

## 🔑 Environment Setup

Create `.env` (development) and `.env.production` (production) files with:

```env
# Server Configuration
HOST=0.0.0.0
APP_VERSION=1.0.0
APP_PORT=8085
API_MODE=debug
GIN_MODE=debug

# Security
JWT_SECRET=your_jwt_secret
INTERNAL_KEY=your_internal_key
ENCRYPT_KEY=your_encryption_key

# Logging
LOGGER_LEVEL=debug
LOGGER_MODE=production
LOGGER_ENCODING=console

# Telegram Integration
TELEGRAM_BOT_KEY=your_telegram_bot_key
TELEGRAM_TELE_LEAD=your_telegram_chat_id
TELEGRAM_DIRECT_LEAD=your_telegram_chat_id

# Database Configuration
POSTGRES_HOST=postgres
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=postgres

# Email Configuration
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_USERNAME=your_email
SMTP_PASSWORD=your_app_password
SMTP_FROM=your_email
SMTP_FROM_NAME=Your Name

# Redis Configuration
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0

# RabbitMQ Configuration
RABBITMQ_HOST=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest

# Cloudinary Configuration
CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret
```

## 🚀 Installation & Setup

1. Clone the repository:
```bash
git clone https://gitlab.com/tantai-smap/authenticate-api.git
cd authenticate-api
```

2. Install dependencies:
```bash
go mod download
```

3. Database setup:
```bash
make container-db      # Start PostgreSQL container
make copy-schema      # Copy initial schema
make migrate-init     # Initialize migrations
make copy-migration   # Copy migration files
make migrate-all      # Run all migrations
```

4. Generate database models:
```bash
make models
```

5. Generate Swagger documentation:
```bash
make swagger
```

## 🏃‍♂️ Running the Application

### Development Mode
```bash
make run-api          # Run API server
make run-consumer     # Run message consumer
```

### Production Mode (Docker)
```bash
docker-compose up -d
```

The API will be available at `http://localhost:8085`

## 📚 API Documentation

Access Swagger documentation at:
```
http://localhost:8085/swagger/index.html
```

## 🛠 Development Tools

### Database Management
```bash
make migrate-create name=migration_name  # Create new migration
make migrate-all                        # Run all migrations
make migrate-down                       # Rollback last migration
```

### Code Generation
```bash
make models           # Generate database models
make swagger          # Update API documentation
make mock            # Generate mocks for testing
```

## 🌐 Docker Network

The application uses an external Docker network named `local-dev_default`. Services:
- API Server: Port 8085
- Message Consumer
- PostgreSQL
- Redis
- RabbitMQ

## 📝 License

This project is licensed under the MIT License.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request
