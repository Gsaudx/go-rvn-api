FROM golang:1.23.6 AS builder

WORKDIR /go-rvn-api

COPY go.mod go.sum ./

RUN go mod download                 

COPY . .

WORKDIR /go-rvn-api/cmd/api

RUN go build -o ../../main .

FROM golang:1.23.6

WORKDIR /app

COPY --from=builder /go-rvn-api/main .
COPY .env .env

# Command to run the executable
CMD ["./main"]