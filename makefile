build: proto-lint
	docker build -t testing --target test .

test: build
	docker run -v $(PWD)/artifacts:/artifacts -v /var/run/docker.sock:/var/run/docker.sock testing
	cd artifacts && curl -s https://codecov.io/bash | bash

proto-lint:
	docker run -v "$(PWD)/docs/protos:/work" uber/prototool:latest prototool lint

proto-generate: proto-lint
	mkdir -p internal/protos
	docker run -v "$(PWD)/docs/protos:/work" -v $(PWD)/pkg/:/output -u `id -u $(USER)`:`id -g $(USER)` uber/prototool:latest prototool generate

test-proto-generate:
	mkdir -p internal/testing/testProto/v1

	protoc --proto_path=docs/testing/proto/ --go_out=paths=source_relative:internal/testing/testProto/ docs/testing/proto/v1/*.proto
	protoc --proto_path=docs/testing/proto/ --go_out=paths=source_relative:internal/testing/testProto/ docs/testing/proto/gen/*.proto

ship:
	docker login --username ${DOCKER_USERNAME} --password ${DOCKER_PASSWORD}
	docker build -t syncromatics/proto-schema-registry:${VERSION} --target final .
	docker push syncromatics/proto-schema-registry:${VERSION}
