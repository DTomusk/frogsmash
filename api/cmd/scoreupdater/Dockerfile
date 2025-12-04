# Stage 1: Build binaries 
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

COPY internal/infrastructure/email/templates ./internal/infrastructure/email/templates

# Build the application, output to bin folder
RUN go build -o /bin/frogsmash-score-updater ./cmd/scoreupdater

# Stage 2: Create the final image
FROM alpine:3.22.2

ENV ENV=production

RUN apk --no-cache add ca-certificates

COPY --from=builder /bin/frogsmash-score-updater /usr/local/bin/
COPY --from=builder /app/internal/infrastructure/email/templates ./internal/infrastructure/email/templates

CMD ["/usr/local/bin/frogsmash-score-updater"]