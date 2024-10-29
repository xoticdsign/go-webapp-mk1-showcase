# go-webapp-mark1-showcase

This project is a server-rendered web application built with the Go programming language, leveraging the GoFiber framework. It implements secure user registration, login, and profile management with JWT-based authentication. This README provides an overview of the application’s structure, setup instructions, and a summary of each key component.

## Features

	•	User Authentication: Registration and login with password hashing using bcrypt and JWT-based session management.
	•	Role-Based Access Control: Access restrictions for profile management and editing.
	•	Server-Rendered HTML: HTML templates served for various pages.
	•	Persistent User Data: Utilizes PostgreSQL with GORM ORM to manage user information.

## Project Structure

go-webapp-mark1-showcase/
├── gobcrypt/                  # Password hashing with bcrypt
│   └── gobcrypt.go            # Encrypt and Decrypt functions
├── gojwt/                     # JWT token management
│   └── gojwt.go               # Token generation and verification
├── gorm/                      # Database and ORM (PostgreSQL with GORM)
│   └── gorm.go                # User model and DB configuration
├── server/
│   ├── handlers/              # HTTP handlers for routes
│   │   └── handlers.go        # Core route handlers
│   ├── middleware/            # Middleware for request handling
│   │   └── auth.go            # Authentication middleware
│   └── templates/             # HTML templates for server-rendered pages
├── .env                       # Environment variables (JWT secret, DB credentials)
├── cert.crt                   # SSL certificate (for HTTPS)
├── cert.key                   # SSL key (for HTTPS)
└── main.go                    # Application entry point

## Core Components

	1.	gobcrypt: Handles password encryption and decryption.
	•	Encrypt: Hashes a plain-text password using bcrypt.
	•	Decrypt: Validates a password against a bcrypt hash.
	2.	gojwt: Manages JWT token generation and verification.
	•	ConfigJWT: Creates a JWT for authenticated sessions.
	•	VerifyJWT: Validates a given JWT.
	3.	gorm: Contains the User struct, PostgreSQL configuration, and functions for database operations.
	•	ConfigPostgreSQL: Initializes a PostgreSQL connection and sets up the user table if needed.
	•	SelectUser: Verifies user credentials during login.
	•	CreateUser: Creates a new user during registration.
	•	UpdateUser: Updates user profile information.
	4.	server/handlers: Contains HTTP route handlers for user interaction.
	•	Login, Registration, Profile: Handlers for rendering pages and processing form submissions.
	•	MainPage: Renders the main landing page.
	•	ProfileEdit: Manages the profile edit page and submits profile changes.
	5.	server/middleware: Contains the Auth middleware for route protection.
	•	Auth: Restricts access to certain pages based on user authentication and role.

## Setup Instructions

### Prerequisites

	•	Go 1.17+
	•	PostgreSQL database
	•	SSL certificate (cert.crt and cert.key)

### Installation

	1.	Clone the Repository:

git clone https://github.com/your-username/go-webapp-mark1-showcase.git
cd go-webapp-mark1-showcase


	2.	Set Environment Variables:
Create a .env file with the following keys:

JWT_SECRET=your_jwt_secret
POSTGRESQL_HOST=your_postgres_host
POSTGRESQL_USER=your_postgres_user
POSTGRESQL_PASSWORD=your_postgres_password
POSTGRESQL_DBNAME=your_db_name
POSTGRESQL_SSLMODE=disable


	3.	Install Dependencies:

go mod download


	4.	Run the Application:

go run main.go


	5.	Access the Application:
Visit https://localhost:6524 in your browser.

Usage

	•	Home (/): Public landing page.
	•	Login (/login): Login form for existing users.
	•	Registration (/registration): User registration form.
	•	Profile (/profile): Displays user profile information; requires authentication.
	•	Profile Edit (/profile/edit): Edit user profile details.

## Routes

Method	Endpoint	Description
GET	/	Main landing page
GET	/login	Login page
POST	/login/login-submit	Submit login credentials
GET	/registration	Registration page
POST	/registration/registration-submit	Submit registration form
GET	/logout	Logout user
GET	/profile	Profile page (Auth required)
GET	/profile/edit	Edit profile page
POST	/profile/edit/edit-submit	Submit profile edits

## Security Features

	•	JWT Authentication: Ensures secure login sessions with JWT.
	•	Bcrypt Password Hashing: Encrypts user passwords before storing them in the database.
	•	HTTPS: Provides secure communication over HTTPS with SSL certificates.

## License

This project is licensed under the MIT License.

This structure will give users clear guidance on setting up, running, and using the app.