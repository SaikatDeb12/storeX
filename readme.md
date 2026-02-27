## Intro
Asset Management System API is a production-ready backend service designed to manage organizational assets, user authentication, role-based access control, and asset lifecycle operations. It provides secure JWT-based authentication, session management, and fine-grained role authorization for handling enterprise-level asset workflows.

## Features

- User Registration and Login: Secure authentication with hashed passwords.

- JWT-Based Authentication: Token-based authentication with session tracking.

- Session Management: Server-side session validation for enhanced security.

- Role-Based Access Control (RBAC): Restrict critical asset operations based on user roles.

- User Management: Fetch all users and retrieve user details by ID.

- Asset Creation: Add new assets with structured validation.

- Asset Fetching: Retrieve assets with protected access.

- Asset Updating: Modify asset details securely.

- Asset Assignment: Assign assets to users.

- Asset Service Handling: Mark assets as sent to service.

- Transactional Database Operations: Atomic user and session creation.

- Centralized Error Handling: Structured and consistent API responses.

- Health Check Endpoint: Monitor server availability.

Tech Stack

Backend: Go (Golang)

Router: Chi

Database: PostgreSQL

ORM/Query Layer: sqlx

Authentication: JWT

Architecture: Modular layered architecture (handler, middleware, dbhelper, utils)

API Routes
Base URL
/v1
Public Routes
Health Check
GET /v1/health
Register User
POST /v1/auth/register

Request Body:

{
  "name": "string",
  "email": "string",
  "phoneNumber": "string",
  "role": "string",
  "employment": "string",
  "password": "string"
}

Response:

{
  "message": "user register successfully",
  "token": "jwt_token"
}
Login User
POST /v1/auth/login

Request Body:

{
  "email": "string",
  "password": "string"
}

Response:

{
  "message": "login successfull",
  "token": "jwt_token"
}
Protected Routes (Requires Authorization Header)
Authorization: Bearer <jwt_token>
User Routes
Get All Users
GET /v1/users
Get User By ID
GET /v1/users/{id}
Asset Routes
Fetch Assets
GET /v1/asset
Create Asset
POST /v1/asset
Update Asset
PUT /v1/asset/update/{id}
Assign Asset
PUT /v1/asset/assign
Send Asset to Service
PUT /v1/asset/service/{id}
Logout
POST /v1/auth/logout

Response:

{
  "message": "Logged out successfully"
}
Authentication Flow

User registers or logs in.

Server generates:

User ID

Session ID

JWT Token containing user_id, session_id, and role.

Client stores token.

Token must be included in the Authorization header.

Middleware validates:

Token signature

Session validity

Role permissions

Project Structure
/cmd
/internal
  /handler
  /middleware
  /dbhelper
  /database
  /models
  /utils
Installation

Clone the repository:

git clone https://github.com/SaikatDeb12/storeX.git
cd asset-management-system

Configure environment variables in .env:

DATABASE_URL=
JWT_SECRET=

Run database migrations.

Start the server:

go run cmd/main.go
Usage

Register or login to obtain a JWT token.

Include the token in the Authorization header.

Access protected routes based on assigned roles.

Manage assets securely through the provided endpoints.
