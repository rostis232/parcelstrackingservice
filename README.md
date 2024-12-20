# Parcel Tracking Service

Parcel Tracking Service is a Go-based application that allows users to track parcels using their tracking numbers. The service fetches tracking information from the 4PX tracking website and provides it via a REST API.

## Project Structure

```
cmd/
    api/
        main.go
config/
    config.yml
internal/
    handler/
        handler.go
    parser/
        manager_test.go
        manager.go
        parser_test.go
        parser.go
        scraper_test.go
        scraper.go
    pkg/
        app/
            app.go
    routes/
        routes.go
    service/
        service.go
models/
    models.go
go.mod
go.sum
trackingservice.dockerfile
docker-compose.yml
```

## Getting Started

### Prerequisites

- Go 1.23.3 or later
- Docker
- Docker Compose

### Configuration

The application configuration is stored in config.yml:

```yml
port: "8181"
max-queue-count: 10
timeout: "30s"
```

### Building and Running

#### Using Docker

1. Build and run the Docker container:

    ```sh
    docker-compose up --build
    ```

2. The service will be available at http://localhost:8181


#### Running Locally

1. Install dependencies:

    ```sh
    go mod download
    ```

2. Build the application:

    ```sh
    go build -o trackingservice ./cmd/api
    ```

3. Run the application:

    ```sh
    ./trackingservice
    ```

## API Endpoints

### Track Parcel

- **URL:** `/track`
- **Method:** `POST`
- **Request Body:**

    ```json
    {
        "tracking_number": "4PX3001521662170CN"
    }
    ```

- **Response:**

    ```json
    {
        "origin_country": "CN",
        "destination_country": "IE",
        "checkpoints": [
            {
                "status": "Delivery information has been provided to An Post",
                "date": "2024-12-18 19:02"
            }
        ]
    }
    ```

## Project Structure

- **cmd/api/main.go:** Entry point of the application.
- **config/config.yml:** Configuration file.
- **models/models.go:** Data models.
- **internal/pkg/app/app.go:** Application setup and initialization.
- **internal/routes/routes.go:** Route definitions.
- **internal/handler/handler.go:** HTTP handler for the tracking endpoint.
- **internal/service/service.go:** Service layer for handling business logic.
- **internal/parser:** Contains the logic for parsing tracking information.

## Running Tests

To run the tests, use the following command:

```sh
go test ./...
```

## License

This project is licensed under the MIT License.