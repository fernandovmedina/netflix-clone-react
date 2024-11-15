# Build stage
FROM golang:1.23.2 as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o main .

# Final stage
FROM debian:bullseye-slim

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]

