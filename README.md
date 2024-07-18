# HTTP Proxy Server

## Overview

This project is an HTTP proxy server that forwards requests to third-party services and returns responses in JSON format. It is built using Go and Gorilla Mux, and includes Docker for containerization.

## Features

- Proxies HTTP requests to third-party services.
- Returns responses in JSON format.
- Validates incoming requests.
- Saves requests and responses locally using `sync.Map`.

## Prerequisites

- Go 1.19 or later
- Docker
- Docker Compose

## Installation

1. Clone the repository:

    ```sh
    git clone <repository-url>
    cd proxy_server
    ```

2. Build and run the Docker container:

    ```sh
    docker-compose up --build
    ```

## Usage

Send a POST request to `http://localhost:8080/api/proxy` with the following JSON body:

```json
{
  "method": "GET",
  "url": "http://google.com",
  "headers": { "Authentication": "Basic bG9naW46cGFzc3dvcmQ=", ... }
}
