FROM golang:1.24-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o api-service cmd/api/main.go
EXPOSE 8083
CMD ["./api-service"]