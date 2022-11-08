# PULL Golang image
FROM golang:1.19.3 AS builder

RUN env
# Using /app directory to store all code
WORKDIR /app

COPY ./app/ .

# Build the app
RUN go build -o scanner

ENTRYPOINT ["./scanner"]
# Leave for later upgrade
#FROM scratch
#COPY --from=builder /app/getweather .
## Start running the app
#CMD ["/getweather"]