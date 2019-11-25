build:
	docker build -t testing --target test .

test: build
	docker run -v $(PWD)/artifacts:/artifacts testing
	cd artifacts && curl -s https://codecov.io/bash | bash

generate-proto:
	mkdir -p internal/testing/testProto/v1

	protoc -I docs/testing/proto/ --go_out=.:internal/testing/testProto/ docs/testing/proto/v1/*.proto
	protoc -I docs/testing/proto/ --go_out=.:internal/testing/testProto/ docs/testing/proto/gen/*.proto