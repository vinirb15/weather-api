## Features

- 


## Installation

To use this project, you need to follow these steps:

1. Clone the repository: `git clone https://github.com/username/repo-name.git`
2. Install the dependencies: `go mod download`
3. Build the application: `go build`
4. Run the application: `./main`

## Docker and Docker Compose

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment. Here are the most common Docker and Docker Compose commands you may want to use:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -p 8080:8080 -e PORT=8080 your-image-name`: Run a container based on the built image. Replace `your-image-name` with the name you used when building the image. You can change the port number if necessary.

If you want to use Docker Compose, follow these commands:

- `docker compose build`: Build the services defined in the `docker-compose.yml` file.
- `docker compose up`: Run the services defined in the `docker-compose.yml` file.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```sh
docker-compose down
```

For more information on Docker and Docker Compose, refer to the official documentation:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Used Tools

This project uses the following tools:

- [Golang](https://golang.org/) for backend development
- [Go-Gin](https://github.com/gin-gonic/gin) for route management
- [GoORM](https://gorm.io/) for database communication
- [SQLite](https://www.sqlite.org/index.html) as the database
- [Swagger](https://swagger.io/) for API documentation and testing

## Usage

After the API is running, you can use the Swagger UI to interact with the endpoints for searching, creating, editing, and deleting job opportunities. The API can be accessed at `http://localhost:$PORT/swagger/index.html`.

Default $PORT if not provided=8080.
