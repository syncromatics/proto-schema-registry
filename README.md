# proto-schema-registry

[![Action](https://github.com/syncromatics/proto-schema-registry/workflows/build/badge.svg)](https://github.com/syncromatics/proto-schema-registry/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/syncromatics/proto-schema-registry)](https://goreportcard.com/report/github.com/syncromatics/proto-schema-registry)

Proto Schema Registry provides a service to store proto schemas in kafka. It also supports breaking change detection and versioning.
Protobuf generated code is both backwards and forward compatible but this service provides a means for other services that don't have
the generated protobuf objects the ability to decode and display messages.

## Getting Started

### Installing

#### Docker
```docker run -e KAFKA_BROKER=broker:9092 -e PORT=443 syncromatics/proto-schema-registry:v0.7.1```

| Environment Variable      | Description                                                                                    |
|---------------------------|------------------------------------------------------------------------------------------------|
| KAFKA_BROKER              | A seed kafka broker used to discover the cluster that schema will be stored in. (required)     |
| PORT                      | The port the grpc service will be hosted at. (required)                                        |
| REPLICATION_FACTOR        | The replication factor of the schema topic. (defaults to 3)                                    |
| SECONDS_TO_WAIT_FOR_KAFKA | The seconds to wait for kakfa to be available before the service panics (defaults to 30).      |

### Usage

Check out the [Protobuf Definition](docs/protos/proto/schema/registry/v1/registry_api.proto).

Once the service is up and running use the `RegisterSchema` grpc method to register a schema. The schema must be gzipped in the message.
You can retrieve schemas by id using the `GetSchema` grpc method. Before the schema is sent it should be extracted from the proto message
by using the [ExtractSchema](pkg/protobuf/extractor.go) function available in `pkg`. This will flatten out the object definition and
generate one file.

## Development

### Requirements

This project requires at least [Go](https://golang.org/dl/) 1.13.1.

### Building

To build and test the project simply run ```make test```.

## Built With

* [proto](https://github.com/emicklei/proto) - Decoder for protobuf files.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [proto](https://github.com/emicklei/proto) for decoding proto files without the need for protoc
* [Visual Studio Code](https://code.visualstudio.com/) for just being an all around great editor
