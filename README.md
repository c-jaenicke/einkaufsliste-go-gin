# einkaufsliste-go-gin

A simple shopping list application that saves entries in a PostgreSQL database. It can serve those entries using a REST API or simple HTML pages and forms.

## Frontend

A basic vue.js frontend exists [c-jaenicke/einkaufsliste-vue](https://github.com/c-jaenicke/einkaufsliste-vue).

Possibly more to come.

## Authentication

**There is no built-in authentication!**

You should put the site behind [Authelia (authelia.com)](https://www.authelia.com/) or another authentication service to
protect it.

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
  # optional, in case you dont have a postgresql database already runnings
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

## Language / Translation

Everything but the code is in german, fork it and translate it if you want to.

## Database

The site requires a PostgreSQL database to save entries.

<sup><sub>Because im lazy i only did PostgreSQL.</sub></sup>

## Icons

All used icons are taken from [https://feathericons.com/](https://feathericons.com/)
