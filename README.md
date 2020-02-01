# simplex-server
Server implementing simplex messaging protocol in Go

## Deploying to Heroku

```sh
$ heroku create
$ git push heroku master
$ heroku open
```

or

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)


## Development

### DB migrations

Install [golang-migrate/migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

Run migrations locally:

```bash
migrate -path=migrations -database postgres://localhost:5432/postgres?sslmode=disable up
```


### Running server

```bash
PORT=8080 simplex-server
```


### Running tests

To run the tests, the server should be running.

If server runs on `localhost:8080`:

```bash
go test
```

If server runs on another URI:

```bash
go test -server=http://example.com
```
