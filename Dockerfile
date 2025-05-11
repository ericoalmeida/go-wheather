FROM golang:1.22 AS builder

WORKDIR /server

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/api

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /server

COPY --from=builder /server/app .

RUN adduser -D -u 10001 appuser
USER appuser

EXPOSE 8080

ENTRYPOINT ["./app"]
