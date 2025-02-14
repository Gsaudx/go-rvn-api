FROM golang:1.23.6 AS builder

WORKDIR /go-tasks-api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /go-tasks-api/cmd/api

RUN go build -o ../../main .

FROM golang:1.23.6

WORKDIR /app

COPY --from=builder /go-tasks-api/main .
COPY .env .env

# Command to run the executable
CMD ["./main"]