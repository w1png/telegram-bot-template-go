FROM golang:1.21.5-bullseye AS builder
WORKDIR /app
COPY . .
RUN apt-get update
RUN apt-get install -y build-essential libsqlite3-dev
RUN go mod download
ENV CGO_ENABLED=1
RUN go build -tags "linux" -o bot .

FROM debian:bullseye
WORKDIR /app
COPY --from=builder /app/bot /app/bot
COPY --from=builder /app/image[s] /app
ENTRYPOINT ["/app/bot"]
