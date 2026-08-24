FROM golang:1.27-alpine3.24 AS builder
WORKDIR /app

COPY . .

WORKDIR /app/services/eloc_auth

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/services/eloc_auth/main .
COPY --from=builder /app/services/eloc_auth/app.env .

EXPOSE 8081

CMD [ "/app/main" ]