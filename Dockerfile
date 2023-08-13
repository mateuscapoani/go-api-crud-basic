FROM golang:1.20.6 as build
WORKDIR /api
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o aplicacao ./cmd/api/main.go

FROM scratch
COPY --from=build /api/aplicacao /aplicacao
COPY --from=build /api/.env /.env
ENTRYPOINT ["/aplicacao"]
