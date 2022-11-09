# PULL Golang image
FROM golang:1.19.3 AS builder

# Using /app directory to store all code
WORKDIR /app

COPY ./app/ .

# Build the app
RUN env CGO_ENABLED=0 go build -tags netgo -ldflags '-w -s -extldflags "-static"' -o getweather

# New image only with our needed binary
FROM alpine
COPY --from=builder /app/getweather .

## Start running the app
ENTRYPOINT ["./getweather"]