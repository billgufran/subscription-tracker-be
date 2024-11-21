# Subscription Tracker API

A robust REST API built with Go for managing and tracking subscription services. This application helps users keep track of their recurring subscriptions, payment methods, and billing cycles.

## Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Configuration](#configuration)
  - [Running the Application](#running-the-application)
- [Project Structure](#project-structure)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
  - [Categories](#categories)
  - [Billing Cycles](#billing-cycles)
  - [Payment Methods](#payment-methods)
  - [Subscriptions](#subscriptions)
- [Database](#database)

## Features

- **User Authentication**: Secure registration and login using JWT.
- **Category Management**: Create, read, update, and delete subscription categories.
- **Billing Cycle Management**: Manage different billing cycles like monthly, yearly, etc.
- **Payment Method Management**: Handle various payment methods such as credit cards, bank accounts, and digital wallets.
- **Subscription Tracking**: Track active subscriptions, next billing dates, and reminders.
- **Default Data Seeding**: Automatically seeds default categories, currencies, and billing cycles.

## Technology Stack

- **Language**: Go
- **Framework**: Gin
- **ORM**: GORM
- **Database**: PostgreSQL
- **Authentication**: JWT (JSON Web Tokens)
- **Environment Management**: Godotenv
- **Unique ID Generation**: ULID

## Getting Started

### Prerequisites

- **Go**: Version 1.23.2 or later
- **PostgreSQL**: Installed and running

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/yourusername/subscription-tracker.git
   cd subscription-tracker
   ```

2. **Install Dependencies**

   Ensure you have Go installed. Then, download the necessary Go modules:

   ```bash
   go mod download
   ```

### Configuration

1. **Environment Variables**

   Create a `.env` file in the root directory of the project. Populate it with variables form `.env.example`

   - If using Railway or another platform that provides a `DATABASE_URL`, you can set that instead of individual DB parameters.

2. **Database Setup**

   The application uses GORM's AutoMigrate feature to handle database migrations. When you run the application, it will automatically create the necessary tables and seed default data if they don't exist.

### Running the Application

Start the server using the following command:

```bash
go run cmd/api/*.go
```

The server will start on `http://localhost:8080` (or the port specified in your `.env` file).

## Project Structure

```plaintext
subscription-tracker/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── auth/
│   │   ├── jwt.go
│   │   └── password.go
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── database.go
│   ├── handlers/
│   ├── middleware/
│   │   └── auth_middleware.go
│   ├── models/
│   │   └── types.go
│   ├── repository/
│   ├── server/
│   │   ├── server.go
│   │   └── routes.go
│   ├── services/
│   └── utils/
│       ├── errors.go
│       ├── http.go
│       └── user_service.go
├── go.mod
├── go.sum
└── README.md
```

### Root Level
- `cmd/` - Contains the main application entry points
  - `api/main.go` - The main application bootstrap file that initializes and starts the server

- `internal/` - Private application code that can't be imported by other projects
  - `auth/` - Authentication related code
    - `jwt.go` - JWT token generation and validation
    - `password.go` - Password hashing and verification

  - `config/` - Configuration management
    - `config.go` - Loads and manages environment variables and app configuration

  - `database/` - Database initialization and management
    - `database.go` - Handles database connection, migrations, and seeding

  - `handlers/` - HTTP request handlers (controllers)
    - Contains route handlers that process incoming HTTP requests

  - `middleware/` - HTTP middleware components
    - `auth_middleware.go` - Authentication middleware for protected routes
    - Has a planned CORS middleware (see reference to middleware/cors_middleware.go)

  - `models/` - Database models and types
    - `types.go` - Common types used across the application
    - Contains structs that represent database tables

  - `repository/` - Database access layer
    - Contains interfaces and implementations for database operations
    - Follows repository pattern for data access

  - `server/` - HTTP server setup
    - `server.go` - Server initialization and configuration
    - `routes.go` - Route definitions and setup

  - `services/` - Business logic layer
    - Contains business logic that sits between handlers and repositories
    - Handles validation and complex operations

  - `utils/` - Shared utilities
    - `errors.go` - Custom error types and error handling
    - `http.go` - HTTP response helpers
    - `user_service.go` - User-related utility functions

### Project Files
- `go.mod` - Go module definition and dependencies
- `go.sum` - Dependency checksums for reproducible builds
- `.env.example` - Environment variables example values
- `README.md` - Project documentation

This structure follows clean architecture principles, separating concerns into distinct layers (handlers → services → repositories → database) and keeping the code organized and maintainable.

## API Endpoints

### Authentication

- **Register**

  ```http
  POST /api/v1/auth/register
  ```

  **Request Body:**

  ```json
  {
    "email": "user@example.com",
    "password": "securepassword",
    "name": "John Doe"
  }
  ```

- **Login**

  ```http
  POST /api/v1/auth/login
  ```

  **Request Body:**

  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```

### Categories

- **Get All Categories**

  ```http
  GET /api/v1/categories
  ```

- **Create Category**

  ```http
  POST /api/v1/categories
  ```

  **Request Body:**

  ```json
  {
    "name": "Streaming"
  }
  ```

- **Update Category**

  ```http
  PUT /api/v1/categories/:id
  ```

  **Request Body:**

  ```json
  {
    "name": "New Category Name"
  }
  ```

- **Delete Category**

  ```http
  DELETE /api/v1/categories/:id
  ```

### Billing Cycles

- **Get All Billing Cycles**

  ```http
  GET /api/v1/billing-cycles
  ```

- **Create Billing Cycle**

  ```http
  POST /api/v1/billing-cycles
  ```

  **Request Body:**

  ```json
  {
    "name": "Monthly",
    "days": 30
  }
  ```

- **Update Billing Cycle**

  ```http
  PUT /api/v1/billing-cycles/:id
  ```

  **Request Body:**

  ```json
  {
    "name": "Updated Cycle Name",
    "days": 31
  }
  ```

- **Delete Billing Cycle**

  ```http
  DELETE /api/v1/billing-cycles/:id
  ```

### Payment Methods

- **Get All Payment Methods**

  ```http
  GET /api/v1/payment-methods
  ```

- **Create Payment Method**

  ```http
  POST /api/v1/payment-methods
  ```

  **Request Body:**

  ```json
  {
    "name": "Visa Ending in 1234",
    "type": "credit_card",
    "lastFour": "1234"
  }
  ```

- **Update Payment Method**

  ```http
  PUT /api/v1/payment-methods/:id
  ```

  **Request Body:**

  ```json
  {
    "name": "Updated Payment Method Name",
    "type": "debit_card",
    "lastFour": "5678"
  }
  ```

- **Delete Payment Method**

  ```http
  DELETE /api/v1/payment-methods/:id
  ```

### Subscriptions

- **Get All Subscriptions**

  ```http
  GET /api/v1/subscriptions
  ```

- **Get Subscription by ID**

  ```http
  GET /api/v1/subscriptions/:id
  ```

- **Create Subscription**

  ```http
  POST /api/v1/subscriptions
  ```

  **Request Body:**

  ```json
  {
    "name": "Netflix",
    "description": "Streaming Service",
    "amount": 9.99,
    "categoryId": "your-category-ulid",
    "currencyId": "your-currency-ulid",
    "billingCycleId": "your-billing-cycle-ulid",
    "paymentMethodId": "your-payment-method-ulid",
    "nextBillingDate": "2024-05-01T00:00:00Z",
    "reminderDays": 5
  }
  ```

- **Update Subscription**

  ```http
  PUT /api/v1/subscriptions/:id
  ```

  **Request Body:**

  ```json
  {
    "name": "Updated Subscription Name",
    "description": "Updated Description",
    "amount": 14.99,
    "categoryId": "updated-category-ulid",
    "currencyId": "updated-currency-ulid",
    "billingCycleId": "updated-billing-cycle-ulid",
    "paymentMethodId": "updated-payment-method-ulid",
    "nextBillingDate": "2024-06-01T00:00:00Z",
    "reminderDays": 7,
    "active": true
  }
  ```

- **Delete Subscription**

  ```http
  DELETE /api/v1/subscriptions/:id
  ```

## Database

Subscription Tracker uses PostgreSQL as its primary database. The connection details are managed via environment variables.

### Migrations and Seeding

On server startup, GORM's `AutoMigrate` feature will automatically create the necessary tables if they do not exist. Additionally, the application seeds default categories, currencies, and billing cycles to ensure the system has essential data to function correctly.

### Default Data

- **Categories**: Includes system-defined categories like Streaming, Gaming, Music, etc.
- **Currencies**: Common currencies such as USD, EUR, GBP, IDR, etc.
- **Billing Cycles**: Standard billing cycles like Weekly, Monthly, Quarterly, etc.