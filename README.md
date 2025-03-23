# URL SHORTNER

Simple URL shortener using **Golang**, [PostgreSQL](https://www.postgresql.org/), [Fiber](https://gofiber.io/) and [GORM](https://gorm.io/)

## Features

- Uniqueness: the same long url will always generate same short url
- Stats Tracking: you can go to `GET /stats/<short-url>` to know how many clicks happened to short url
- Scalable: We can use kubernetes to scale this system with load balancing. `Dockerfile` is also present.
- Custom Length: You can set `LENGTH` param in `.env` to set the Length of shortened url. Default is 8
- Custom Alias

## Setup

Rename `.env.exmaple` to `.env`

Run database:

You can use docker or install `postgre` on your machine and create a database and fill the `.env` file accordingly

```sh
docker run --name dbname \
-e POSTGRES_USER=$DB_USERNAME -e POSTGRES_PASSWORD=$DB_PASSWORD \
-p $DB_PORT:$DB_PORT -d postgres:14
```

Run application:

Either run `go run main.go` or follow below to run on localmachine

```bash
go build -o shortner
./shortener
```

## REST API

### Create Go-short

`POST /shorten`

```sh
curl -i -H 'Content-Type: application/json' \
-d '{"redirect":"http://github.com"}' \
-X POST http://localhost:3000/shorten
```

For using custom Alias

```sh
curl -i -H 'Content-Type: application/json' \
-d '{"redirect":"http://github.com"}' \
-X POST http://localhost:3000/shorten?alias=customname
```

### Redirect

`GET /<short-url>`

### Stats

`GET /stats/<short-url>`
