FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY internal/infrastructure/email/templates ./internal/infrastructure/email/templates

RUN go build -o /bin/worker ./cmd/worker

FROM alpine:3.22.2

ENV ENV=production

COPY --from=builder /bin/worker /bin/worker
COPY --from=builder /app/internal/infrastructure/email/templates ./internal/infrastructure/email/templates

CMD ["/bin/worker"]