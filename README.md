# Backend challenge for Klever

**Warning:** This project is a work in progress. There are still tasks to be completed before it's fully functional. 

- Some unit and integration tests still need to be created. 
- A Makefile will be added for build and deployment automation.
- Instructions for starting the application via a Docker container will be provided.

## Table of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)

## Introduction

Restful API created with [Go](https://go.dev/) and the [Gin](https://gin-gonic.com/) framework that implements some functions of a bitcoin wallet. 
## Getting Started

**Prerequisites:**

- Make sure you have [Go](https://go.dev/), version 1.21.4 or higher, installed on your device;
- Create a valid .env file following the .env.example template that is in the root of the project;
- Ensure that no application is listening on port :8080.

**Preparing the environment:**

Before running locally, run the command below to ensure that all dependencies will be installed:

```shell
go mod tidy
```

**Running the API without build:**

At the root of the project, run the command 

```shell
go run ./main.go
```

**Running the API with build (on Linux or Mac):**

At the root of the project, run the command 

```shell
go build -o api ./main.go && ./api
```


## API Endpoints

The base url will be: `http://localhost:8080/api/v1`

**Health** [GET]: `/health`

_Response:_

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

Try with curl: `curl -svX GET 'http://localhost:8080/api/v1/health'`

External API response time is calculated by making three simultaneous calls to the external api using goroutines.

**Details** [GET]: `/details/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n`

_Response:_

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

Try with curl: `curl -svX GET 'http://localhost:8080/api/v1/details/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n'`

**Balance** [GET]: `/balance/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n`

_Response:_

```json
{
  "confirmed": "12845845",
  "unconfirmed": "0"
}
```

Try with curl: `curl -svX GET 'http://http://localhost:8080/api/v1/balance/19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n'`

**Send** [POST]: `/send`

_Response:_

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

Try with curl: `curl -svX POST "http://localhost:8080/api/v1/send" -H 'Content-Type: application/json' -d '{"address": "19SH3YrkrpWXKtCoMXWfoVpmUF1ZHAi24n", "amount": "1208053"}'`

**Transaction** [GET]: `/tx/3654d26660dcc05d4cfb25a1641a1e61f06dfeb38ee2279bdb049d018f1830ab`

_Response:_

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


## Testing

The tests were created with built-in Go packages and the [testify](https://github.com/stretchr/testify) toolkit.

To run the tests implemented so far, at the root of the project run:

```shell
go test ./...
```

For a more verbose result:

```shell
go test -v ./...
```
