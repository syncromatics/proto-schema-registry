package service

import (
	"context"

	v1 "github.com/syncromatics/proto-schema-registry/pkg/proto/schema/registry/v1"
)

// Storage is the storage interface for schema storage
type Storage interface {
	GetSchema(ctx context.Context, id int64) (schema []byte, ok bool, err error)
	RegisterSchema(ctx context.Context, topic string, schema []byte) (id int64, errors []string, ok bool, err error)
}

// Service is the implementation of the schema service
type Service struct {
	storage Storage
}

// GetSchema returns a schema by id
func (s *Service) GetSchema(ctx context.Context, request *v1.GetSchemaRequest) (*v1.GetSchemaResponse, error) {
	schema, ok, err := s.storage.GetSchema(ctx, request.Id)
	if err != nil {
		return nil, err
	}

	if !ok {
		return &v1.GetSchemaResponse{
			Exists: false,
		}, nil
	}

	return &v1.GetSchemaResponse{
		Exists: true,
		Schema: schema,
	}, nil
}

// RegisterSchema registers a schema or returns the id if it already exists
func (s *Service) RegisterSchema(ctx context.Context, request *v1.RegisterSchemaRequest) (*v1.RegisterSchemaResponse, error) {
	id, schemaErrors, ok, err := s.storage.RegisterSchema(ctx, request.Topic, request.Schema)
	if err != nil {
		return nil, err
	}

	if !ok {
		return &v1.RegisterSchemaResponse{
			Response: &v1.RegisterSchemaResponse_ResponseError{
				ResponseError: &v1.RegisterSchemaError{
					Errors: schemaErrors,
				},
			},
		}, nil
	}

	return &v1.RegisterSchemaResponse{
		Response: &v1.RegisterSchemaResponse_ResponseSuccess{
			ResponseSuccess: &v1.RegisterSchemaSuccess{
				Id: id,
			},
		},
	}, nil
}

// NewService creates a new schema service
func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}
