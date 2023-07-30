# DOCKERFILE FOR BUILDING REST IMAGE
# BUILD USING `docker build -f dockerfile-rest . -t einkaufsliste-rest:latest`
## BUILD
FROM golang:1.20 as build-stage

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY cmd/enting/ ./
COPY pkg/ ./pkg/
COPY ent/ ./ent/
RUN CGO_ENABLED=0 GOOS=linux go build -o /enting

## DEPLOY
FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /enting /app/enting
COPY .env ./

ENV GIN_MODE release
EXPOSE 8080

CMD [ "/app/enting" ]
