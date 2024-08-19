ARG CONFIG_PATH

FROM golang:1.20.3-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o ./bin/matchmaker cmd/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/bin/matchmaker .
COPY config.env .

CMD ["./matchmaker"]