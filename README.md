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
