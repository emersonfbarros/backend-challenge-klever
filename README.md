# Backend challenge for Klever

**Warning:** This project is a work in progress. There are tasks to complete so that all challenge requirements are met. But the application itself is already complete and working. 

- Some integration tests still need to be created. 

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)

## Introduction

Restful API created with [Go](https://go.dev/) and the [Gin](https://gin-gonic.com/) framework that implements some functions of a bitcoin wallet. 
## Getting Started

### Prerequisites:

- Make sure you have [Go](https://go.dev/), version 1.21.4 or higher, installed on your device;
- Create a valid .env file following the .env.example template that is in the root of the project;
- Ensure that no application is listening on port :8080.

### Preparing the environment

Before running locally, run the command below to ensure that all dependencies will be installed:

```shell
go mod download
```

### Running locally

**Running the API without build:**

At the root of the project, run the command:

```shell
go run ./main.go
```

### Using Makefile commands

The project includes a Makefile to help you manage common tasks more easily. Here's a list of the available commands and a brief description of what they do:

- `make run`: Run the application in gin debug mode.
- `make build`: Build the application and create an executable file named `gowallet`.
- `make release`: Run the application in gin release mode.
- `make test`: Run tests for all packages in the project.
- `make test-verbose`: Run tests for all packages in the project in verbose mode.
- `make test-coverage`: Runs tests for all packages and shows coverage.
- `make clean`: Remove the `gowallet` executable.

To use these commands, simply type `make` followed by the desired command in your terminal. For example:

```shell
make run
```

### Docker and Docker compose

This project includes a `Dockerfile` and `docker-compose.yml` file for easy containerization and deployment.

**Running application with Docker:**

At the root of the project:

- `docker build -t your-image-name .`: Build a Docker image for the project. Replace `your-image-name` with a name for your image.
- `docker run -d -p 8080:8080 --name your-container-name your-image-name:latest`: Run a container based on the built image in detached mode. Replace `your-container-name` with a name for container and replace `your-image-name` with the name you used when building the image.
- `docker stop your-container-name`: Stop the container running the application.
- `docker start your-container-name`: Restarts the previously initialized container.

**Using Docker compose:**

If you want to use Docker Compose, follow these commands in the root of the project:

- `docker compose up -d`: Build and run the services defined in the `docker-compose.yml` file in detached mode.

To stop and remove containers, networks, and volumes defined in the `docker-compose.yml` file, run:

```shell
docker-compose down
```

## API Endpoints

The base url will be: `http://localhost:8080/api/v1`

**Health** [GET]: `/health`

Try with curl: `curl -svX GET 'http://localhost:8080/api/v1/health'`

_Response:_

External API response time is calculated by making three simultaneous calls to the external api using goroutines.

HTTP code: `200 OK`
```json
{
  "status": "healthy",
  "timestamp": "2023-12-07T11:58:38-03:00",
  "uptime": "0D 0H 15M 9S",
  "externalApi": {
    "status": "Ok",
    "responseTime": "1169ms"
  }
}
```

**Details** [GET]: `/details/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n`

_Possible responses:_

Try with curl: `curl -svX GET 'http://localhost:8080/api/v1/details/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n'`

HTTP code: `200 OK`
```json
{
  "address": "19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n",
  "balance": "12845845",
  "totalTx": 642,
  "balanceCalc": {
    "confirmed": "12845845",
    "unconfirmed": "0"
  },
  "total": {
    "sent": "176043318",
    "received": "188889163"
  }
}
```

HTTP code: `404 Not Found`
```json
{
    "message": "Adress <invalid_address> not found"
}
```

HTTP code: `500 Internal Server Error`
```json
{
    "message": "Internal server error"
}
```

HTTP code: `500 Bad Gateway`
```json
{
    "message": "Failed to request external resource"
}
```

**Balance** [GET]: `/balance/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n`

Try with curl: `curl -svX GET 'http://localhost:8080/api/v1/balance/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n'`

_Response:_

HTTP code: `200 OK`
```json
{
  "confirmed": "12845845",
  "unconfirmed": "0"
}
```

HTTP code: `404 Not Found`
```json
{
    "message": "Adress <invalid_address> not found"
}
```

HTTP code: `500 Internal Server Error`
```json
{
    "message": "Internal server error"
}
```

HTTP code: `500 Bad Gateway`
```json
{
    "message": "Failed to request external resource"
}
```

**Send** [POST]: `/send`

Request body:

```json
{
  "address": "<string> - Bitcoin address",
  "amount": "<string> - The amount to send (in Satoshis). Note that this should be a string representation of a number"
}
```

Try with curl: `curl -svX POST "http://localhost:8080/api/v1/send" -H 'Content-Type: application/json' -d '{"address": "19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n", "amount": "1208053"}'`

_Response:_

HTTP code: `200 OK`
```json
{
  "utxos": [
    {
      "txid": "3672b48dcb1494a30ad94b2cd092cc380d6bb475b86ca547655a65c0c27941e5",
      "amount": "1146600"
    },
    {
      "txid": "4148e2e46000c3a988ae2b7687a40b2bab7f29fb822ecc22912566d7b74330a4",
      "amount": "1100593"
    }
  ]
}
```

HTTP Code `400 Bad Request`
```json
{
    "message": "Error message will depend on the specific issue with the request body"
}
```

HTTP code: `404 Not Found`
```json
{
    "message": "Adress <invalid_address> not found"
}
```

HTTP code: `500 Internal Server Error`
```json
{
    "message": "Internal server error"
}
```

HTTP code: `500 Bad Gateway`
```json
{
    "message": "Failed to request external resource"
}
```

**Transaction** [GET]: `/tx/3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab`

_Response:_

Try with curl `curl -svX GET 'http://localhost:8080/api/v1/tx/3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab'`

HTTP code: `200 OK`
```json
{
  "addresses": [
    {
      "address": "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
      "value": "484817655"
    },
    {
      "address": "36iYTpBFVZPbcyUs8pj3BtutZXzN6HPNA6",
      "value": "623579"
    },
    {
      "address": "bc1qyzxdu4px4jy8gwhcj82zpv7qzhvc0fvumgnh0r",
      "value": "422126277"
    }
    ...
  ],
  "block": 675674,
  "txID": "3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab"
}
```

HTTP code: `404 Not Found`
```json
{
    "message": "Adress <invalid_transaction_id> not found"
}
```

HTTP code: `500 Internal Server Error`
```json
{
    "message": "Internal server error"
}
```

HTTP code: `500 Bad Gateway`
```json
{
    "message": "Failed to request external resource"
}
```


## Testing

The tests were created with built-in Go packages and the [testify](https://github.com/stretchr/testify) toolkit.

To run the tests implemented so far you can use the makefile commands mentioned above or, at the root of the project, run:

```shell
go test ./...
```

For a more verbose result:

```shell
go test -v ./...
```

To see test coverage:

```shell
go test -v ./... -coverprofile=cover.out
```
