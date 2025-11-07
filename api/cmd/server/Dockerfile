# Stage 1: Build binaries 
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

COPY db/migrations ./db/migrations

# Build the application
RUN go build -o /bin/frogsmash-api ./cmd/server

# Stage 2: Create the final image
FROM alpine:3.22.2

ENV ENV=production

RUN apk --no-cache add ca-certificates

COPY --from=builder /bin/frogsmash-api /usr/local/bin/
COPY --from=builder /app/db/migrations ./db/migrations

CMD ["/usr/local/bin/frogsmash-api"]