# GoLang MVC with PostgreSQL

## Overview
This project demonstrates a web application built using GoLang, the Model-View-Controller (MVC) architecture, and PostgreSQL as the database. It provides a scalable, maintainable, and well-structured foundation for building robust web applications.

## Features
- MVC architecture for separation of concerns.
- PostgreSQL integration for data persistence.
- RESTful API design.
- Clean code with proper package structuring.
- Authentication using JWT (JSON Web Tokens).
- Basic CRUD operations with Bearer token access.
- Docker support for easy deployment.

## Prerequisites
- Go 1.16+ installed on your system.
- PostgreSQL 10+ installed and running.
- Docker (optional, for containerized deployment).

## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/VirendraArekar/go-postgres-boilerplate.git
    cd go-postgres-boilerplate
    git init crud
    ```

2. **Set up environment variables**:
    Create a `.env` file in the root directory to store environment variables like database credentials:
    ```
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=yourusername
    DB_PASSWORD=yourpassword
    DB_NAME=yourdbname
    ```

3. **Install dependencies**:
    ```bash
    go mod download
    ```

4. **Set up the database**:
    Create the database and run migrations (if any):
    ```bash
    psql -U yourusername -d postgres -c "CREATE DATABASE yourdbname;"
    ```

4. **Apply Migration**:
    Create the database and run migrations (if any):
    ```bash
    go run database/migrate.go -action=up  # Apply All migrations

    go run database/migrate.go -action=down -steps=1 # Roll Back Migrations

    go run database/migrate.go -action=drop  # Drop the database if it exists.
    ```

5. **Run the application**:
    ```bash
    go run main.go
    ```

6. **Access the application**:
    Open your browser and go to `http://localhost:8080`.

## Endpoints

### Example auth routes:
- `POST /auth/register`: Register user.
- `POST /auth/login`: Login user.

### Example users CRUD routes:
- `GET /api/users`: Fetch all users.
- `GET /api/user/{id}`: Get a single user by ID.
- `PUT /api/user/{id}`: Update an existing user by ID.
- `DELETE /api/user/{id}`: Delete an user by ID.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for any bugs or feature requests.

## Author
**Virendra Arekar**  
Feel free to reach out at [virendra.arekar@gmail.com](mailto:virendra.arekar@gmail.com)
