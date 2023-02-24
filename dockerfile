## BUILD
FROM golang:1.19-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd/v1/ ./
COPY pkg/ ./pkg/

RUN go build -o /shopping

## DEPLOY
FROM alpine
WORKDIR /app
COPY --from=builder /shopping /app/shopping
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/assets /app/assets
COPY --from=builder /app/images /app/images

ENV PORT 8080
ENV GIN_MODE release
EXPOSE 8080

CMD [ "/app/shopping" ]
