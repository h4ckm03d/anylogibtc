# anylogbtc

## Constraint

- Authorization will be skipped focus on the task deliver 2 endpoints stated on the pdf

## Prerequisites
- Docker is required for simple setup env using docker compose
- Makefile to simplify command execution
- Golang 1.19
- [jq](https://stedolan.github.io/jq/) ( optional for better formatting output)

## Development

### Setup development environment

- Run `docker compose up -d` to setup environment
- Migration script automatically run when running docker compose up. All migration on the `init.sql`


- Install dependencies
```cmd
go mod tidy
go install -tags sqlite github.com/gobuffalo/pop/v6/soda@latest
```
Or simply run `make dep` command.


### How to run and test

Run command `go run cmd/main.go` to run application.

If you are using custom database url please specify `DATABASE_URL` on environment variable. By default it will use setting on docker compose

Run unit test
```
go test -mod=readonly -v ./... -covermode=count -coverprofile=profile.out && go tool cover -func=profile.out
```

or 

```
make test
```
example output
```
❯ make test
go test -mod=readonly -v ./... -covermode=count -coverprofile=profile.out && go tool cover -func=profile.out
=== RUN   TestHealthz
--- PASS: TestHealthz (0.00s)
=== RUN   TestSaveTransaction
--- PASS: TestSaveTransaction (0.00s)
=== RUN   TestSaveFailTransaction
--- PASS: TestSaveFailTransaction (0.00s)
=== RUN   TestGetHistory
--- PASS: TestGetHistory (0.00s)
PASS
coverage: 57.4% of statements
ok      anylogibtc/api/handler  0.003s  coverage: 57.4% of statements
?       anylogibtc/cmd  [no test files]
?       anylogibtc/dto  [no test files]
?       anylogibtc/entity       [no test files]
?       anylogibtc/repository   [no test files]
?       anylogibtc/repository/pg        [no test files]
?       anylogibtc/repository/repositoryfakes   [no test files]
?       anylogibtc/services/transaction [no test files]
?       anylogibtc/services/transaction/transactionfakes    [no test files]
anylogibtc/api/handler/server.go:28:            NewEchoServer0.0%
anylogibtc/api/handler/server.go:36:            SetupRoutes 0.0%
anylogibtc/api/handler/server.go:49:            Run         0.0%
anylogibtc/api/handler/server.go:73:            Healthz     100.0%
anylogibtc/api/handler/transaction.go:18:       NewTransactionHandler        100.0%
anylogibtc/api/handler/transaction.go:25:       Save        100.0%
anylogibtc/api/handler/transaction.go:46:       History     100.0%
```

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
├── repository
│   ├── pg
│   │   └── transaction.go
│   ├── repository.go
│   ├── transaction.go
│   └── walletfakes
│       └── fake_transaction_repository.go
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

| field      | Type      | Description                  |
| :--------- | :-------- | :--------------------------- |
| `datetime` | `string`  | **Required**. ISODATE string |
| `amount`   | `float64` | **Required**. btc amount     |

Examples:
```bash
curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2019-10-05T14:45:05+07:00","amount": 10}'
```

Results:
```
```

#### Get history

```http
  POST /v1/wallets/history
```
Request body:
| Field           | Type     | Description                   |
| :-------------- | :------- | :---------------------------- |
| `startDatetime` | `string` | **Required**. Start date time |
| `endDatetime`   | `string` | **Required**. End date time   |

The result history will be in UTC format
Examples:
```bash
curl -X POST "localhost:3000/v1/wallets/history" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"startDatetime":"2022-12-05T02:45:05+07:00","endDatetime":"2022-12-05T07:45:05+07:00"}'

[
  {
    "datetime": "2022-12-04T14:00:00Z",
    "amount": "14"
  }
]
```
## Notes

- Timescaledb used because I need continuous aggregations to speed up the calculation on each hour. For better understanding I put the migration script to create hypertable and materialized view.
```sql
-- create hypertable on table transactions, before create this id and datetime must be indexed together as Primary key
SELECT
  *
FROM
  create_hypertable(
    'transactions',
    'datetime',
    chunk_time_interval => INTERVAL '1 hour'
  );

--
CREATE MATERIALIZED VIEW transaction_hourly
WITH
  (timescaledb.continuous) AS
SELECT
  time_bucket(INTERVAL '1 hour', datetime) AS hour,
  SUM(amount) as total
FROM
  transactions
GROUP BY
  hour;
```

## Examples run
```bash
curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:05+07:00","amount": 10}'
{"message":"data created successfully"}

curl -X POST "localhost:3000/v1/wallets/history" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"startDatetime":"2022-12-05T02:45:05+07:00","endDatetime":"2022-12-05T07:45:05+07:00"}'
[{"datetime":"2022-12-05T00:00:00Z","amount":"10"}]

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:06+07:00","amount": 2}' 
{"message":"data created successfully"}

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:05+07:00","amount": 2}' 
{"message":"data created successfully"}

curl -X POST "localhost:3000/v1/wallets/history" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"startDatetime":"2022-12-05T02:45:05+07:00","endDatetime":"2022-12-05T07:45:05+07:00"}'
[{"datetime":"2022-12-05T00:00:00Z","amount":"14"}]

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T07:45:05+07:00","amount": 10}'
{"message":"data created successfully"}

curl -X POST localhost:3000/v1/wallets  \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"datetime": "2022-12-05T08:45:05+07:00","amount": 1.32}'
{"message":"data created successfully"}

curl -X POST "localhost:3000/v1/wallets/history" \
-H "Accept: application/json" \
-H "Content-Type: application/json" \
-d '{"startDatetime":"2022-12-05T02:45:05+07:00","endDatetime":"2022-12-05T09:45:05+07:00"}'
[{"datetime":"2022-12-05T00:00:00Z","amount":"24"},{"datetime":"2022-12-05T01:00:00Z","amount":"25.32"}]
```

## Authors

- [@h4ckm03d](https://www.github.com/h4ckm03d)


