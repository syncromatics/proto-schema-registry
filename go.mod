module github.com/syncromatics/proto-schema-registry

go 1.13

require (
	docker.io/go-docker v1.0.0
	github.com/Microsoft/go-winio v0.4.14 // indirect
	github.com/Shopify/sarama v1.24.1
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v0.7.3-0.20190817195342-4760db040282 // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/emicklei/proto v1.8.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/google/uuid v1.1.1
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/klauspost/cpuid v1.2.3 // indirect
	github.com/opencontainers/go-digest v1.0.0-rc1 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.4.0
	go.uber.org/zap v1.13.0
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4
	google.golang.org/genproto v0.0.0-20200413115906-b5235f65be36 // indirect
	google.golang.org/grpc v1.28.1
	gotest.tools v2.2.0+incompatible
)

replace github.com/docker/docker/internal/testutil => gotest.tools/v3 v3.0.0
