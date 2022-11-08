# PULL Golang image
FROM golang:1.19.3 AS builder

RUN env
# Using /app directory to store all code
WORKDIR /app

COPY ./app/ .

# No need for go dependencies to install
RUN go mod download

# Build the app
RUN go build -o getweather

ENTRYPOINT ["./getweather"]
# Leave for later upgrade
#FROM scratch
#COPY --from=builder /app/getweather .
## Start running the app
#CMD ["/getweather"]