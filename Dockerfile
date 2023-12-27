FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/IlyaZayats/managus
COPY . .

RUN go build -o ./bin/managus ./cmd/managus

FROM alpine:latest AS runner

COPY --from=builder /go/src/github.com/IlyaZayats/managus/bin/managus /app/managus

RUN apk -U --no-cache add bash ca-certificates \
    && chmod +x /app/managus

WORKDIR /app
ENTRYPOINT ["/app/managus"]
