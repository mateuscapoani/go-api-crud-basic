FROM golang:1.20.6
WORKDIR /api
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o aplicacao ./cmd/api/main.go
ENTRYPOINT ["/api/aplicacao"]
