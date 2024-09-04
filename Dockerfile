FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go test ./...
RUN go build -o /app/main ./cmd/notifier/main.go

CMD ["/app/main"]
