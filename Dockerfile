# Stage 1: Build the app
FROM golang:alpine AS builder

WORKDIR /app

# Copy the Go module manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download && go mod verify

# Copy the source code
COPY . .

# Build Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main ./cmd

# Stage 2: Create a minimal image

FROM alpine:latest

# Set the working app
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy .env file
COPY --from=builder /app/.env .

# Copy entrypoint sh
COPY --from=builder /app/scripts/entrypoint.sh .

# Wait-for-it requires bash, which alpine doesn't ship with by default.
# Use wait-for instead
ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for ./entrypoint.sh

# Command to run the binary
ENTRYPOINT [ "sh", "./entrypoint.sh" ]
