# Getting Started

## Prerequisites

Install the postgres version of `sqlc` and `golang-migrate`.

```bash
$ go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
...
$ go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Create new schema

```bash
$ migrate create -ext sql -dir ./migrations -seq schema
...
```

## Starting your dockerised database

```bash
$ docker run -e POSTGRES_USER=user -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=fullstackgo -p 5432:5432 postgres:11.10-alpine
...
PostgreSQL init process complete; ready for start up.

2022-01-24 08:13:00.658 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
2022-01-24 08:13:00.658 UTC [1] LOG:  listening on IPv6 address "::", port 5432
2022-01-24 08:13:00.662 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
2022-01-24 08:13:00.676 UTC [49] LOG:  database system was shut down at 2022-01-24 08:13:00 UTC
2022-01-24 08:13:00.683 UTC [1] LOG:  database system is ready to accept connections

```

## Migrate database to latest

```bash
$ migrate -path ./migrations -database postgres://user:mysecretpassword@0.0.0.0:5432/fullstackgo?sslmode=disable up
1/u schema (61.189882ms)
```
