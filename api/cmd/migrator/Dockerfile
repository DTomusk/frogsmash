# Stage 1: Build the migrator binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /bin/migrator ./cmd/migrator

# Stage 2: Create the final image
FROM alpine:3.22.2

ENV ENV=production

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/db/migrations ./db/migrations
COPY --from=builder /bin/migrator /usr/local/bin/

CMD ["/usr/local/bin/migrator"]
