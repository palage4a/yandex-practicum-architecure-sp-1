FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o example.com/temperature-api

FROM alpine:latest
RUN apk --no-cache add ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/example.com/temperature-api .
EXPOSE 8080
CMD ["./temperature-api"]