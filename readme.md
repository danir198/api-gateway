# API Gateway

This repository contains an API Gateway for an e-commerce application. The API Gateway routes requests to the Inventory Service and provides SSL/TLS support for secure communication.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Setup](#setup)
- [Generating SSL Certificates](#generating-ssl-certificates)
- [Running the Application](#running-the-application)
- [Environment Variables](#environment-variables)
- [Project Structure](#project-structure)

## Prerequisites

- Docker
- Docker Compose
- Go 1.20 or later

## Setup

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/api-gateway.git
    cd api-gateway
    ```

2. Install Go dependencies:

    ```sh
    go mod tidy
    ```

## Generating SSL Certificates

To enable HTTPS, you need to generate SSL certificates. You can use OpenSSL to create self-signed certificates for testing purposes.

1. Create a directory for the certificates:

    ```sh
    mkdir tls
    ```

2. Generate the certificate and key files:

    ```sh
    openssl req -x509 -newkey rsa:4096 -keyout tls/key.pem -out tls/cert.pem -days 365 -nodes -subj "/CN=localhost"
    ```

## Running the Application

You can run the application using Docker Compose.

1. Build and start the services:

    ```sh
    docker-compose up --build
    ```

2. The API Gateway will be available at `https://localhost:8002`.

## Environment Variables

The following environment variables are used in the `docker-compose.yml` file:

- `MONGO_URI`: MongoDB connection URI.
- `DATABASE_NAME`: Name of the MongoDB database.
- `SERVER_ADDRESS`: Address for the Inventory Service.
- `ENABLE_AUTH`: Enable authentication for the Inventory Service.
- `API_USERNAME`: Comma-separated list of API usernames.
- `API_PASSWORD`: Password for the API users.
- `JWT_SECRET`: Secret key for JWT authentication.
- `JWT_EXPIRATION`: JWT token expiration time in seconds.

## Project Structure

. ├── docker-compose.yml ├── Dockerfile ├── gateway.go ├── go.mod ├── go.sum ├── main.go └── tls/ ├── cert.pem └── key.pem


hardcoded-credentials Embedding credentials in source code risks unauthorized access

- [`docker-compose.yml`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdani.rosyadi_1%2FDocuments%2Fupworks%2Fecommerce%2Fapi-gateways%2Fdocker-compose.yml%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22ffdcf805-415a-49f8-87b9-41e2874e5fea%22%5D "/Users/dani.rosyadi_1/Documents/upworks/ecommerce/api-gateways/docker-compose.yml"): Docker Compose configuration file.
- [`Dockerfile`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdani.rosyadi_1%2FDocuments%2Fupworks%2Fecommerce%2Fapi-gateways%2FDockerfile%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22ffdcf805-415a-49f8-87b9-41e2874e5fea%22%5D "/Users/dani.rosyadi_1/Documents/upworks/ecommerce/api-gateways/Dockerfile"): Dockerfile for building the API Gateway image.
- [`gateway.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdani.rosyadi_1%2FDocuments%2Fupworks%2Fecommerce%2Fapi-gateways%2Fgateway.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22ffdcf805-415a-49f8-87b9-41e2874e5fea%22%5D "/Users/dani.rosyadi_1/Documents/upworks/ecommerce/api-gateways/gateway.go"): Contains the implementation of the API Gateway.
- [`go.mod`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdani.rosyadi_1%2FDocuments%2Fupworks%2Fecommerce%2Fapi-gateways%2Fgo.mod%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22ffdcf805-415a-49f8-87b9-41e2874e5fea%22%5D "/Users/dani.rosyadi_1/Documents/upworks/ecommerce/api-gateways/go.mod"): Go module file.
- [`go.sum`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdani.rosyadi_1%2FDocuments%2Fupworks%2Fecommerce%2Fapi-gateways%2Fgo.sum%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22ffdcf805-415a-49f8-87b9-41e2874e5fea%22%5D "/Users/dani.rosyadi_1/Documents/upworks/ecommerce/api-gateways/go.sum"): Go dependencies file.
- [`main.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdani.rosyadi_1%2FDocuments%2Fupworks%2Fecommerce%2Fapi-gateways%2Fmain.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22ffdcf805-415a-49f8-87b9-41e2874e5fea%22%5D "/Users/dani.rosyadi_1/Documents/upworks/ecommerce/api-gateways/main.go"): Entry point for the API Gateway.
- [`tls/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fdani.rosyadi_1%2FDocuments%2Fupworks%2Fecommerce%2Fapi-gateways%2Ftls%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%2C%22ffdcf805-415a-49f8-87b9-41e2874e5fea%22%5D "/Users/dani.rosyadi_1/Documents/upworks/ecommerce/api-gateways/tls/"): Directory containing SSL certificate and key files.

## License

This project is licensed under the MIT License.