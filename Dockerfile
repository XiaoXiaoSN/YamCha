FROM golang:1.15-alpine as builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN apk add --update git ca-certificates 
RUN go build -o yamcha .

# pull the binary file and service work really in the layer
FROM alpine:latest

WORKDIR /srv/yamcha
RUN touch config.yml
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/yamcha /srv/yamcha/yamcha
COPY --from=builder /app/configs/config-build.yml /srv/yamcha/configs/config.yml

ENTRYPOINT ["./yamcha", "line"]
