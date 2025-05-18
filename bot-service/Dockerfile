FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o bot-service cmd/bot/main.go
CMD ["./bot-service"]