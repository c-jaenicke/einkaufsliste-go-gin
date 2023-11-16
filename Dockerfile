# DOCKERFILE FOR BUILDING REST IMAGE
# BUILD USING `docker build -f dockerfile-rest . -t einkaufsliste-rest:latest`
## BUILD
FROM golang:1.20 as build-stage

WORKDIR /app

# copy all files in directory
COPY . .
# get and update packages
RUN go mod tidy
# build go app
RUN CGO_ENABLED=0 GOOS=linux go build -o /enting cmd/enting/main.go

## DEPLOY
FROM alpine:latest

WORKDIR /app

COPY --from=build-stage /enting /app/enting
COPY .env ./

ENV GIN_MODE release
EXPOSE 8080

CMD [ "/app/enting" ]
