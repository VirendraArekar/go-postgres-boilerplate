# GoLang MVC with PostgreSQL

## Overview
This project demonstrates a web application built using GoLang, the Model-View-Controller (MVC) architecture, and PostgreSQL as the database. It provides a scalable, maintainable, and well-structured foundation for building robust web applications.

## Features
- MVC architecture for separation of concerns.
- PostgreSQL integration for data persistence.
- RESTful API design.
- Clean code with proper package structuring.
- Basic CRUD operations.
- Docker support for easy deployment.

## Prerequisites
- Go 1.16+ installed on your system.
- PostgreSQL 10+ installed and running.
- Docker (optional, for containerized deployment).

## Project Structure
├── config # Configuration files (DB credentials, etc.) ├── controllers # Handles HTTP requests and responses ├── models # Database models and interactions ├── routes # Routing logic for the application ├── views # HTML templates (if applicable) ├── database # Database migration and seeding ├── utils # Utility functions and helpers ├── main.go # Entry point of the application └── README.md # Project documentation


## Installation

1. **Clone the repository**:
    ```bash
    git clone https://github.com/VirendraArekar/go-postgres-boilerplate.git
    cd go-postgres-boilerplate
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
    go run database/migrate.go  # Apply migrations
    ```

5. **Run the application**:
    ```bash
    go run main.go
    ```

6. **Access the application**:
    Open your browser and go to `http://localhost:8080`.

## Endpoints

### Example CRUD routes:
- `GET /items`: Fetch all items.
- `POST /items`: Create a new item.
- `GET /items/{id}`: Get a single item by ID.
- `PUT /items/{id}`: Update an existing item by ID.
- `DELETE /items/{id}`: Delete an item by ID.

## Docker Support
If you want to run the application in a Docker container:
1. Build the Docker image:
    ```bash
    docker build -t golang-mvc-postgres .
    ```

2. Run the container:
    ```bash
    docker run -p 8080:8080 golang-mvc-postgres
    ```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing
Contributions are welcome! Please submit a pull request or open an issue for any bugs or feature requests.

## Author
**Your Name**  
Feel free to reach out at [your.email@example.com](mailto:your.email@example.com)
