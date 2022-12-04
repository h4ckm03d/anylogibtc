# anylogbtc

## Constraint

- Authorization will be skipped focus on the task deliver 2 endpoints stated on the pdf

## Prerequisites
- Docker is required for simple setup env using docker compose
- Makefile to simplify command execution
- Golang 1.19

## Development

### Setup development environment

- Run `docker compose up -d` to setup environment 

- Install dependencies
```cmd
go mod tidy
go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```
Or simply run `make dep` command.

- Run migration database using `soda`
```
	soda migrate -e development
```
or run `make migrate` command

## Directory structure

```
├── LICENSE
├── Makefile
├── README.md
├── api
│   ├── README.md
│   └── handler
│       ├── server.go
│       ├── server_test.go
│       ├── transaction.go
│       └── transaction_test.go
├── cmd
│   └── main.go
├── database.yml
├── docker-compose.yml
├── domain
│   └── wallet
│       ├── pg
│       │   └── transaction.go
│       ├── repository.go
│       ├── transaction.go
│       └── walletfakes
│           └── fake_transaction_repository.go
├── dto
├── entity
├── migrations
├── services
│   └── transaction
│       ├── transaction.go
│       └── transactionfakes
│           └── fake_transaction_service.go
└── tools
    └── tools.go
```

- `api`: implementation API handler, setup routes
- `domain`: location of the repository and its dependencies
- `dto`: all data transfer objects in this directory
- `entity`: all database entity
- `services`: contains all business logic application
- `tools`: This file imports packages that are used when running go generate, or used during the development process but not otherwise depended on by built code.
- `cmd`: entry point application through `cmd/main.go`

## API Reference

#### Get all items

```http
  POST /v1/wallets
```

Request body:

| field       | Type      | Description                  |
| :---------- | :-------- | :--------------------------- |
| `datetime`  | `string`  | **Required**. ISODATE string |
| `amount`    | `float64` | **Required**. btc amount     |

```bash
curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2019-10-05T14:45:05+07:00","amount": 10}'
```

#### Get history

```http
  GET /v1/wallets
```

| Parameter       | Type     | Description                       |
| :-------------- | :------- | :-------------------------------- |
| `startDatetime` | `string` | **Required**. Start date time     |
| `endDatetime`   | `string` | **Required**. End date time       |


## Authors

- [@h4ckm03d](https://www.github.com/h4ckm03d)


