FROM golang:1.19 AS builder

RUN apt-get -y update && apt-get -y upgrade && apt-get install -y sqlite3 libsqlite3-dev

WORKDIR /unit-test-example/

# use modules for deps
COPY go.mod go.sum /unit-test-example/
RUN go mod download

COPY . .

RUN sqlite3 /unit-test-example/tennis.db < setup.sql
RUN CGO_ENABLED=0 GOOS=linux go build -o ./example .

FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates

COPY --from=builder /unit-test-example/example /root/
COPY --from=builder /unit-test-example/tennis.db /root/tennis.db

EXPOSE 8080

CMD ["./example"]
