### ***Installation***



## **Requirements**

Before installing, ensure you have the following dependencies installed on your system:
Go (version 1.20 or later)
Docker
A SQL database manager (e.g., PostgreSQL, MySQL, or SQLite)
Postman or an API client like Insomnia for testing API requests (Optional but recommended) 







## **Setup Steps**



# Clone the repository

git clone https://github.com/robertobouses/easy-football-tycoon.git
cd easy-football-tycoon



# Configure environment variables

Copy the example environment file:
cp .env.config .env
Edit .env and set up the necessary database credentials.



# Start the database and services using Docker

docker-compose up -d



# Run database migrations

go run cmd/migrate.go



# Start the application

go run main.go



# Optional: Running in a Docker Container

If you prefer running the app in Docker:

docker build -t easy-football-tycoon .
docker run --env-file .env -p 5432:5432 easy-football-tycoon



