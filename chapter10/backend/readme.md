# Moving to API First

The supporting code for Chapter 6 of Becoming a Full Stack Go Developer

## Getting Started

To get started you can use an existing running database or follow along and use a dockerised
instance of postgres.

### Docker / Database

We use SQLc and the golang-migrate packages to generate our model code and to handle our
migrations as well as to create new ones. We need to install that first.

```bash
$ go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
$ go get -u -tags 'postgres' github.com/golang-migrate/migrate/cli
...
```

You can now run your migration.

```bash
$ docker run -e POSTGRES_USER=local -e POSTGRES_PASSWORD=asecurepassword -e POSTGRES_DB=fullstackdb -p 5432:5432 postgres:11.10-alpine
$ migrate -path ./migrations -database "postgres://local:asecurepassword@localhost:5432/fullstackdb?sslmode=disable" up
1/u schema (48.903685ms)
```

### Build and run

You can now re-generate the sqlc bindings and run your code.

```bash
$ go generate
$ go build
$ ./chapter6
...
```

### Creating Migrations

To create additional migrations you can run the following command to generate them in sequence.

```bash
$ migrate create -ext sql -dir ./migrations -seq foo
/.../Becoming-a-Full-Stack-Go-Developer/chapter6/migrations/000001_foo.up.sql
/.../Becoming-a-Full-Stack-Go-Developer/chapter6/migrations/000001_foo.down.sql
```

## Usage

Some example queries to check the endpoints

```bash
# Login
curl -H 'Content-Type: application/json' 0.0.0.0:9002/login -d '{"username":"user@user","password":"password"}'  -v

# Get a list of previous workouts and their sets
curl -X GET -H 'Content-Type: application/json' 0.0.0.0:9002/workout --cookie 'session-name=MTY0NTI2MTczNXxEdi1CQkFFQ180SUFBUkFCRUFBQVJQLUNBQUlHYzNSeWFXNW5EQk1BRVhWelpYSkJkWFJvWlc1MGFXTmhkR1ZrQkdKdmIyd0NBZ0FCQm5OMGNtbHVad3dJQUFaMWMyVnlTVVFGYVc1ME5qUUVBZ0FDfMwSOVjl_-nwIrsRVE1b5Q2ss-kd_RyObfoO-HlrVP0j;'

# Add a new workout (with no entries)
curl -X POST -H 'Content-Type: application/json' 0.0.0.0:9002/workout --cookie 'session-name=MTY0NTI2MTczNXxEdi1CQkFFQ180SUFBUkFCRUFBQVJQLUNBQUlHYzNSeWFXNW5EQk1BRVhWelpYSkJkWFJvWlc1MGFXTmhkR1ZrQkdKdmIyd0NBZ0FCQm5OMGNtbHVad3dJQUFaMWMyVnlTVVFGYVc1ME5qUUVBZ0FDfMwSOVjl_-nwIrsRVE1b5Q2ss-kd_RyObfoO-HlrVP0j;'


## Delete a workout
curl -X DELETE -H 'Content-Type: application/json' 0.0.0.0:9002/workout/1 --cookie 'session-name=MTY0NTI2MTczNXxEdi1CQkFFQ180SUFBUkFCRUFBQVJQLUNBQUlHYzNSeWFXNW5EQk1BRVhWelpYSkJkWFJvWlc1MGFXTmhkR1ZrQkdKdmIyd0NBZ0FCQm5OMGNtbHVad3dJQUFaMWMyVnlTVVFGYVc1ME5qUUVBZ0FDfMwSOVjl_-nwIrsRVE1b5Q2ss-kd_RyObfoO-HlrVP0j;'

## Add a set to a workout
curl -X POST -H 'Content-Type: application/json' 0.0.0.0:9002/workout/5 --cookie 'session-name=MTY0NTI2MTczNXxEdi1CQkFFQ180SUFBUkFCRUFBQVJQLUNBQUlHYzNSeWFXNW5EQk1BRVhWelpYSkJkWFJvWlc1MGFXTmhkR1ZrQkdKdmIyd0NBZ0FCQm5OMGNtbHVad3dJQUFaMWMyVnlTVVFGYVc1ME5qUUVBZ0FDfMwSOVjl_-nwIrsRVE1b5Q2ss-kd_RyObfoO-HlrVP0j;' -d '{"exercise_name": "Barbell Rows", "weight":700'}

...
```
