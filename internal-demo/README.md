# Internal User Authentication and Authorization Demo

This project, part of a larger repository, demonstrates how to implement user authentication and authorization internally. It's located in the `internal-demo` folder and serves as an example of creating a robust user management system without heavily relying on third-party providers such as Auth0.

## Overview

The `internal-demo` project showcases various features and best practices in building an internal user management system. It includes user registration, login, password encryption, JWT token management, and more. This approach gives you complete control over user data and authentication flows.

## Features

- User Registration and Login
- Password Hashing and Secure Storage
- JWT (JSON Web Tokens) for Authentication
- Role-Based Access Control (RBAC)
- Integration with MongoDB for data persistence
- CORS Configuration
- Basic frontend UI for interaction

## Technologies Used

- **Backend**: Go (Gin-Gonic Framework)
- **Database**: MongoDB
- **Frontend**: React (with Vite)
- **Authentication**: JWT, bcrypt for password hashing

## Getting Started

### Prerequisites

- Go (version 1.15 or later)
- Node.js and npm
- Docker and Docker Compose (for containerization and easy setup)
- MongoDB instance (local or remote)

### Installation and Setup

1. **Clone the Repository**: Clone the larger repository and navigate to the `internal-demo` folder.
``` bash
git clone https://github.com/lkubicek1/react-go-auth.git
cd internal-demo
```

2. **Set Up the Backend**:
   - Navigate to the backend directory.
   - Copy the `.env.example` file to `.env` and fill in the necessary environment variables.
   - Run the Go server.

``` bash
cd api
cp .env.example .env
# Edit .env with your environment variables
go run main.go
```

3. **Set Up the Frontend**:
   - Navigate to the frontend directory.
   - Install dependencies and run the development server.

``` bash
cd ui
npm install
npm run dev
```

4. **Using Docker Compose** (Optional):
   - Ensure Docker and Docker Compose are installed.
   - Run the following command at the root of `internal-demo` to start all services:

``` bash
docker-compose up --build
```

### Usage

- Access the frontend at `http://localhost:3000` (or the configured port).
- Use the UI to register a new user, log in, and access protected resources.

## License

This project is licensed under the [Apache License 2.0](LICENSE).
