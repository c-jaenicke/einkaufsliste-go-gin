# einkaufsliste-go-gin

A simple shopping list application written in Go, using a PostgreSQL database to store all entries.
Serving a REST-API for the frontend.

## Frontend

A basic vue.js frontend exists [c-jaenicke/einkaufsliste-vue](https://github.com/c-jaenicke/einkaufsliste-vue).

## Authentication

**There is no built-in authentication!**

This doesn't have any authentication included! If you want yours to be secure, put something like Authelia or Authentik
in front ot it.

Everyone that has access to the site can change the entries!

## Docker

You can host the webapp using the included `docker-compose`. Be aware, you need to change some lines to your setup!

Fields marked with `<TEXT>` need to be adjusted!

```yaml
version: "3.8"

services:
  app:
    container_name: shopping-list
    image: shopping:latest
    restart: unless-stopped
    # expose ports of container, but don't bind to host port, useful for reverse proxy
    #expose:
    #  - 8080
    # bind container port to host port, format: HOST:CONTAINER
    ports:
      - 8080:8080
    environment:
      - POSTGRES_URL="postgresql://<POSTGRES_USER>:<POSTGRES_PASSWORD>@172.22.0.2:5432/<POSTGRES_DB>"

  # postgresql database
  # optional, in case you dont have a postgresql database already running
  db:
    container_name: postgres-test
    image: postgres:latest
    restart: unless-stopped
    hostname: postgres-test
    environment:
      - POSTGRES_PASSWORD=<SET PASSWORD HERE>
      - POSTGRES_USER=<SET USER HERE>
      - POSTGRES_DB=<SET DATABASE NAME HERE>
    volumes:
      - postgres-test-volume:/var/lib/postgresql/data

# persistent volume for saving database
volumes:
  postgres-test-volume:
```

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
