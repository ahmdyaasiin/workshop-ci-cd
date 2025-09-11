FROM golang:1.24-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/workshop-ci-cd cmd/api/main.go

FROM alpine:latest

RUN apk add --no-cache ca-certificates curl tzdata && update-ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/workshop-ci-cd .
COPY --from=builder /app/db/migration /app/db/migration

ARG APP_PORT=8001
EXPOSE ${APP_PORT}

ENTRYPOINT ["./workshop-ci-cd"]