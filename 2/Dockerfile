# PULL Golang image
FROM golang:1.19.3 AS builder

# Using /app directory to store all code
WORKDIR /app
# Coping all of our code to the builder image
COPY ./app/ .

# Build the app
# Disable CGO, using netgo tag for net package, and statically linking everythig
RUN env CGO_ENABLED=0 go build -tags netgo -ldflags '-w -s -extldflags "-static"' -o scanner

# Using scratch ensures minimal image footprint and risks
FROM scratch AS final
COPY --from=builder /app/scanner .

## Start running the app
ENTRYPOINT ["./scanner"]