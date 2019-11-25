package service

import (
	"context"

	v1 "github.com/syncromatics/proto-schema-registry/internal/protos/proto/schema/registry/v1"

	"google.golang.org/grpc"
)

type service struct{}

func (s *service) GetSchema(ctx context.Context, request *v1.GetSchemaRequest) (*v1.GetSchemaResponse, error) {
	return nil, nil
}

func (s *service) RegisterSchema(ctx context.Context, request *v1.RegisterSchemaRequest) (*v1.RegisterSchemaResponse, error) {
	return nil, nil
}

// RegisterService will register the registry service with the grpc server
func RegisterService(server *grpc.Server) {
	service := &service{}
	v1.RegisterRegistryAPIServer(server, service)
}
