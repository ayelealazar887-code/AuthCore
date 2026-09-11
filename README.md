# Auth Go JWT 🔐

A backend authentication service built with **Go**, **MySQL**, and **JWT (JSON Web Tokens)**.

This project is being developed to practice and demonstrate how authentication and authorization work in a Go backend application.

## 🚧 Project Status

**In Development**

The goal is to build a complete authentication system with secure user registration, login, JWT-based authentication, refresh tokens, and role-based authorization.

## 🎯 Planned Features

* User registration
* User login
* Password hashing
* JWT access tokens
* Refresh tokens
* Token validation and authentication middleware
* User logout
* Protected routes
* Role-based authorization

  * Admin
  * Employee
* User management
* MySQL database integration
* Database migrations using Goose
* Request validation
* Error handling
* Authentication middleware

## 🛠️ Technologies

* **Go** — Backend language
* **Gin** — HTTP web framework
* **MySQL** — Database
* **JWT** — Authentication
* **Goose** — Database migrations
* **Docker** — MySQL containerization
* **sqlx / database/sql** — Database access

## 📁 Planned Project Structure

```text
Auth/
├── database/
│   └── config/
│       └── db.go
│
├── migrations/
│   └── 00001_create_users_table.sql
│
├── models/
│   └── user.go
│
├── controllers/
│   └── auth.go
│
├── middleware/
│   └── auth.go
│
├── routes/
│   └── routes.go
│
├── services/
│   └── auth.go
│
├── main.go
├── docker-compose.yml
├── go.mod
└── README.md
```

## 🔑 Authentication Flow

The planned authentication flow is:

```text
User
  │
  ▼
Register / Login
  │
  ▼
Validate credentials
  │
  ▼
Hash / Verify password
  │
  ▼
Generate JWT
  │
  ├── Access Token
  │
  └── Refresh Token
  │
  ▼
Client
  │
  ▼
Protected API Request
  │
  ▼
JWT Middleware
  │
  ▼
Validate Token
  │
  ▼
Allow / Reject Request
```

## 👤 User Roles

The application will support different user types:

```text
admin
employee
```

Administrators will have access to administrative functionality, while employees will have access to regular application functionality.

## 🗄️ Database

MySQL will be used to store users and authentication-related information.

The database will run using Docker:

```yaml
services:
  mysql:
    image: mysql:8.4
    container_name: auth_go
    restart: unless-stopped
```

Database migrations will be managed
