# Invoice Management API

A modern invoice management API developed using Gin and GORM.

## Technology Stack

- **Framework**: Gin Web Framework
- **ORM**: GORM (PostgreSQL adapter)
- **Database**: PostgreSQL
- **Configuration Management**: (via `.env` files)
- **Other**: Go 1.21+

## Features

- Invoice CRUD operations
- Dynamic filtering and search
- Pagination
- Quick setup with seed data
- Modern REST API design

## Installation

### Prerequisites

- Go 1.21+
- PostgreSQL 14+
- Git

```bash
# Clone the repository
git clone [git@github.com:logrenant/invoice-app-api.git]
cd invoice-api

# Install dependencies
go mod download
```

## Database Setup

**Create database and user in PostgreSQL:**

- CREATE DATABASE invoices;
- CREATE USER gowit_user WITH PASSWORD 'securepassword123';
- GRANT ALL PRIVILEGES ON DATABASE invoices TO gowit_user;

**Create .env file:**

- DB_HOST=localhost
- DB_PORT=5432
- DB_USER=gowit_user
- DB_PASSWORD=securepassword123
- DB_NAME=invoices

## Configuration

Copy .env.example to .env in the project root directory and update the values.

## Running the Application

**Start the application**
`go run cmd/main.go`

**Load seed data**
`go run cmd/seed/main.go`

## API Documentation

**Core Endpoints**

- **GET** `/invoices` _List invoices_
- **POST** `/invoices` _Create new invoice_
- **GET** `/invoices/{id}` _Get invoice details_
- **PUT** `/invoices/{id}` _Update invoice_
- **DELETE** `/invoices/{id}` _Delete invoice_

### Example Requests

**List Invoices**

```bash
curl -X GET "http://localhost:8080/invoices?page=1&pageSize=10&search=SSP"
```

**Create Invoice**

```bash
curl -X POST "http://localhost:8080/invoices" \
-H "Content-Type: application/json" \
-d '{
  "service_name": "SSP",
  "invoice_number": "INV-2023-001",
  "amount": 1999.99,
  "status": "pending"
}'
```

## Project Structure

```bash
invoice-api/
├── cmd/
│   ├── main.go          # Main application entrypoint
│   ├── seed/            # Seed data scripts
│       └── main.go
├── internal/
│   ├── config/          # Configuration management
│   ├── db/              # PostgreSQL connection
│   ├── handlers/        # HTTP handlers
│   ├── models/          # Database models
│   ├── repositories/    # Database operations
│   └── services/        # Business logic layer
├── .env                 # Environment variables
├── go.mod               # Dependencies
└── go.sum
```
