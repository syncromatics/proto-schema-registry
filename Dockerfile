FROM golang:1.13.1 as build

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go vet ./...

RUN go get -u golang.org/x/lint/golint

RUN golint -set_exit_status ./...

RUN go build -o proto-schema-registry ./cmd/proto-schema-registry

#testing
FROM build as test

CMD go test -race -coverprofile=/artifacts/coverage.txt -covermode=atomic ./...

# final image
FROM ubuntu:18.04 as final

WORKDIR /app

COPY --from=0 /build/proto-schema-registry /app/

CMD ["/app/proto-schema-registry"]
