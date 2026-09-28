FROM golang:1.27-alpine3.24 AS builder
WORKDIR /app

COPY . .

WORKDIR /app/services/eloc_product

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.1/migrate.linux-amd64.tar.gz | tar xvz

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/services/eloc_product/main .
COPY --from=builder /app/services/eloc_product/migrate ./migrate
COPY --from=builder /app/services/eloc_product/app.env .
COPY services/eloc_product/db/migration ./db/migration

COPY services/eloc_auth/start.sh .

EXPOSE 8081

RUN chmod +x /app/start.sh

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]