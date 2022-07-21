FROM golang:1.18.1-alpine
COPY . /build
WORKDIR /build
CMD go run main.go