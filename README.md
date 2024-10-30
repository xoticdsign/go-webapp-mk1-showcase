# go-webapp-mk1-showcase

This project is a server-rendered web application built with the Go programming language, leveraging the GoFiber framework. It implements secure user registration, login, and profile management with JWT-based authentication. It runs on WLAN and can be accesses from any device connected to your WI-FI (check ["Installation"](#installation))! This README provides an overview of the application’s structure, setup instructions, and a summary of each key component.

## Features

- User Authentication: Registration and login with password hashing using bcrypt and JWT.
- Role-Based Access Control: Access restrictions for profile management and editing.
- Server-Rendered HTML: HTML templates served for various pages.
- Persistent User Data: Utilizes PostgreSQL with GORM ORM to manage user information.

## Tech Used

- [gofiber](https://github.com/gofiber/fiber)
- [GORM](https://github.com/go-gorm/gorm)
- [JWT](https://github.com/golang-jwt/jwt)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## Project Structure

Project structure is simplified for showcasing purposes.

```
go-webapp-mark1-showcase/
├── gobcrypt/
│   └── gobcrypt.go
├── gojwt/
│   └── gojwt.go
├── gorm/
│   └── gorm.go
├── server/
│   ├── handlers/
│   │   └── handlers.go
│   ├── middleware/
│   │   └── auth.go
│   └── templates/
│       ├── index.html
│       ├── login.html
│       ├── registration.html
│       ├── profile.html
│       └── profile_edit.html
├── .env
├── cert.crt
├── cert.key
└── main.go
```

## Core Components

1. gobcrypt: Handles password encryption and decryption.
- Encrypt: Hashes a plain-text password using bcrypt.
- Decrypt: Validates a password against a bcrypt hash.
2. gojwt: Manages JWT token generation and verification.
- ConfigJWT: Creates a JWT for authenticated sessions.
- VerifyJWT: Validates a given JWT.
3. gorm: Contains the User struct, PostgreSQL configuration, and functions for database operations.
- ConfigPostgreSQL: Initializes a PostgreSQL connection and sets up the user table if needed.
- SelectUser: Verifies user credentials during login.
- CreateUser: Creates a new user during registration.
- UpdateUser: Updates user profile information.
4. server/handlers: Contains HTTP route handlers for user interaction.
5. server/middleware: Contains the Auth middleware for route protection.
6. server/template: Contains templates for rendering webpages.
7. main.go: Entry point of the application

## Prerequisites

- [Go 1.17+](https://go.dev)
- [PostgreSQL](https://www.postgresql.org/download) installed

## Installation

1. Clone the Repository:

```bash
git clone https://github.com/xoticdsign/go-webapp-mark1-showcase
```

```bash
cd go-webapp-mark1-showcase
```

2. Set Environment Variables:
   
Config your DB connection in .env file in /:

```env
POSTGRESQL_HOST=<your_postgres_host>
POSTGRESQL_USER=<your_postgres_user>
POSTGRESQL_PASSWORD=<your_postgres_password>
POSTGRESQL_DBNAME=<your_db_name>
POSTGRESQL_SSLMODE=<enable/disable>
```

That's all you need to do to get DB running, everything else will be done automatically on initial start up ("users" table will be created, test users will be inserted etc.). You can see test users credentials in the gorm.go, don't try to use credentials from your newly created table, because the passwords are encrypted in there.

3. Run the Application:

```bash
go run main.go
```

4. Access the Application:

Locally

```bash
https://127.0.0.1:6524
```

or on any device connected to your WI-FI.

```bash
https://<ip_of_the_computer_you_started_application_on>:6524
```

## Usage

- Home (/): Public landing page. Can be accessed by anyone.
- Login (/login): Login form. Can be accessed only by visitors.
- Registration (/registration): User registration form. Can be accessed only by visitors.
- Profile (/profile): Displays user profile information. Can be accessed only by users.
- Profile Edit (/profile/edit): Edit user profile details. Can be accessed only by users.

## License

[MIT](https://choosealicense.com/licenses/mit)
