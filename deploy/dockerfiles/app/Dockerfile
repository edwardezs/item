FROM golang:1.22.2 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o APP_EXECUTABLE ./cmd/main.go

FROM alpine:3.17.3
COPY --from=build /app/APP_EXECUTABLE .
CMD ["/APP_EXECUTABLE"]
