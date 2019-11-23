FROM golang:1.13.1 as build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build ./...

RUN go test ./...
