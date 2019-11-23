build:
	docker build .

generate-proto:
	mkdir -p internal/testing/testProto/v1

	protoc -I docs/testing/proto/ --go_out=.:internal/testing/testProto/ docs/testing/proto/v1/*.proto
	protoc -I docs/testing/proto/ --go_out=.:internal/testing/testProto/ docs/testing/proto/gen/*.proto