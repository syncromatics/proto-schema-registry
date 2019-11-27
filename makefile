build: proto-lint
	docker build -t testing --target test .

test: build
	docker run -v $(PWD)/artifacts:/artifacts -v /var/run/docker.sock:/var/run/docker.sock testing
	cd artifacts && curl -s https://codecov.io/bash | bash

proto-lint:
	docker run -v "$(PWD)/docs/protos:/work" uber/prototool:latest prototool lint

proto-generate: proto-lint
	mkdir -p internal/protos
	docker run -v "$(PWD)/docs/protos:/work" -v $(PWD)/internal/protos:/output -u `id -u $(USER)`:`id -g $(USER)` uber/prototool:latest prototool generate

test-proto-generate:
	mkdir -p internal/testing/testProto/v1

	protoc -I docs/testing/proto/ --go_out=.:internal/testing/testProto/ docs/testing/proto/v1/*.proto
	protoc -I docs/testing/proto/ --go_out=.:internal/testing/testProto/ docs/testing/proto/gen/*.proto
