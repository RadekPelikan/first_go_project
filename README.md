# First Go Project

Simple CRUD app

- using React, and standard `net/http` go library

## Running

### Run Migrations

Install tool [golang-migrate/migrate](https://github.com/golang-migrate/migrate) and add to path variables and run

```bash
$ migrate -source file://migrations -database "mysql://root@(localhost:3306)/first_go_project" up
```

### Run Go Server

Clone repository

```bash
$ git clone
```

Download dependencies & build go

```bash
$ go mod tidy
$ go build main.go
```

Run go server

```bash
$ go run ./dist/main.go
```

## Developing

### Create new migration

```bash
$ migrate create -ext sql -dir src/sql/migrations -seq <name>
```
