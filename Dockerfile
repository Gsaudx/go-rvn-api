# Build Stage
FROM golang:1.23.6 as builder

WORKDIR /app
# Cache go modules.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code and build the binary.
COPY . .
RUN CGO_ENABLED=0 go build -o api-go-project ./cmd/api

# Final Stage: minimal runtime environment.
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/api-go-project .

EXPOSE 8080
CMD ["./api-go-project"] 