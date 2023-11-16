# einkaufsliste-go-gin

A simple shopping list application written in Go, using a PostgreSQL database to store all entries.
Serving a REST-API for the frontend.

## Frontend

A basic vue.js frontend exists [c-jaenicke/einkaufsliste-vue](https://github.com/c-jaenicke/einkaufsliste-vue).

A better and more modern looking frontend written using
svelte [c-jaenicke/einkaufsliste-svelte](https://github.com/c-jaenicke/einkaufsliste-svelte).

## Authentication

**There is no built-in authentication!**

This doesn't have any authentication included! If you want yours to be secure, put something like Authelia or Authentik
in front ot it.

Everyone that has access to the site can change the entries!

---

## Docker

```yaml
version: "3.9"

services:
  db:
    container_name: einkaufsliste-db
    hostname: einkaufsliste-db
    image: postgres:15
    restart: unless-stopped
    environment:
      # change values here!
      - POSTGRES_PASSWORD=<SET PASSWORD HERE>
      - POSTGRES_USER=<SET USERNAME HERE>
      - POSTGRES_DB=<SET DB NAME HERE>
    volumes:
      - einkaufsliste-db-volume:/var/lib/postgresql/data
    expose:
      - 5432
    healthcheck:
      test: [ "CMD-SHELL", "pg_isready -d <SET DB NAME HERE> -U <SET USERNAME HERE>" ]
      interval: 10s
      timeout: 5s
      retries: 5

  frontend:
    container_name: einkaufsliste-frontend
    hostname: einkaufsliste-frontend
    image: einkaufsliste-svelte:latest
    restart: unless-stopped
    environment:
      - "ORIGIN=<SET DOMAIN HERE, LOCALHOST:PORT WHEN TESTING>"
    ports:
      - 3000:3000
    expose:
      - 3000

  backend:
    container_name: einkaufsliste-api
    hostname: einkaufsliste-api
    image: einkaufsliste-rest:latest
    restart: unless-stopped
    ports:
      - 8080:8080
    environment:
      # make sure to change values here! need to be the same as above!
      - DSS=host=einkaufsliste-db port=5432 user=<SET USERNAME HERE> dbname=<SET DB NAME HERE> password=<SET PASSWORD HERE> sslmode=disable
      # switch gin mode [debug|test|release]
      - GIN_MODE=release
      # set allowed origins for cors, string of urls seperated by a ","
      - ALLOWED_ORIGINS=http://localhost:3000,http://einkaufsliste-frontend:3000
    expose:
      - 8080
    depends_on:
      db:
        condition: service_healthy

volumes:
  einkaufsliste-db-volume:

networks:
  default:
    driver: bridge
```

### .env

Following string has to be set, either in the `.env` file when building the docker image, or in the docker-compose file.

```env
DSS=host=<db container name> port=5432 user=<username> dbname=<db name> password=<password> sslmode=disable
```

### Test Release Mode Locally

Run using `GIN_MODE=release go run cmd/enting/main.go`

---

## REST-API

### Category

```text
Get     category/all
Post    category/new
Delete  category/:id/delete
Put     (category/:id/update)
```

#### Get

```json
[
  {
    "id": "NUMBER",
    "name": "STRING",
    "color": "STRING(#000000-#ffffff)",
    "edges": {}
  }
]
```

#### Post

```json

{
  "id": "NUMBER",
  "name": "STRING",
  "color": "STRING"
}
```

### Store

```text
Get     store/all
Post    store/new
Delete  store/:id/delete
Put     (store/:id/update)
```

#### Get

```json
[
  {
    "id": "NUMBER",
    "name": "STRING",
    "edges": {}
  }
]
```

#### Post

```json
    {
  "name": "STRING"
}
```

### Item

```text
Get     item/all
Get     item/all?store=store&category=category&=status
Post    item/new
Put     item/:id/update
Delete  item/:id/delete
Patch   item/:id/switch
```

#### Get

```json
[
  {
    "id": "NUMBER",
    "name": "STRING",
    "note": "STRING",
    "amount": "NUMBER",
    "status": "STRING[new|bought]",
    "store_id": "NUMBER",
    "category_id": "NUMBER",
    "edges": {}
  }
]
```

Status can be `new` or `bought`.

#### Post / Put

```json
{
  "name": "STRING",
  "note": "STRING",
  "amount": "NUMBER",
  "store_id": "NUMBER",
  "category_id": "NUMBER"
}
```

### Pet

#### Get

```json
{
	"fed_at": "UNIX TIME AS INT64",
	"amount_fed": "STRING",
	"is_inside": "BOOLEAN"
}
```
